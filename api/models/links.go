package models

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/h00s-go/tiny-link-backend/db/sql"
	"github.com/h00s-go/tiny-link-backend/services"
	"github.com/redis/go-redis/v9"
)

type Links struct {
	services *services.Services
}

func NewLinks(services *services.Services) *Links {
	return &Links{
		services: services,
	}
}

func (ls *Links) FindByShortURI(shortURI string) (*Link, error) {
	return ls.FindByID(IDfromShortURI(shortURI))
}

func (ls *Links) FindByID(id int64) (*Link, error) {
	if l := ls.findInMemstoreByID(id); l != nil {
		return l, nil
	}

	l := &Link{}
	if err := ls.services.DB.Conn.QueryRow(context.Background(), sql.GetLinkByID, id).Scan(&l.ID, &l.URL, &l.CreatedAt); err != nil {
		return nil, err
	}

	go ls.createInMemstore(l)

	return l, nil
}

func (ls *Links) FindByURL(URL string) (*Link, error) {
	if l := ls.findInMemstoreByURL(URL); l != nil {
		return l, nil
	}

	l := &Link{}
	if err := ls.services.DB.Conn.QueryRow(context.Background(), sql.GetLinkByURL, URL).Scan(&l.ID, &l.URL, &l.CreatedAt); err != nil {
		return nil, err
	}

	go ls.createInMemstore(l)

	return l, nil
}

func (ls *Links) Create(URL string) (*Link, error) {
	if l, err := ls.FindByURL(URL); err == nil {
		return l, nil
	}

	l := &Link{URL: URL}
	h := NewHost(l.URL)
	if err := h.IsValid(); err != nil {
		return nil, err
	}

	tx, err := ls.services.DB.Conn.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	if err := tx.QueryRow(context.Background(), sql.CreateLink, l.URL).Scan(&l.ID); err != nil {
		tx.Rollback(context.Background())
		return nil, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	return ls.FindByID(l.ID)
}

// ++++++ Memstore ++++++

func (ls *Links) findInMemstoreByID(id int64) *Link {
	l := &Link{}

	value, err := ls.services.IMDS.Client.Get(context.Background(), "id:"+fmt.Sprint(id)).Result()
	if err == nil {
		if json.Unmarshal([]byte(value), &l) != nil {
			ls.services.Logger.Println("Error while unmarshaling link: ", err)
			return nil
		}
		return l
	} else if err != redis.Nil {
		ls.services.Logger.Println("Error while getting key from memstore: ", err)
	}

	return nil
}

func (ls *Links) findInMemstoreByURL(url string) *Link {
	value, err := ls.services.IMDS.Client.Get(context.Background(), "url:"+url).Result()
	if err == nil {
		id, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			ls.services.Logger.Println("Error converting string to int64:", err)
			return nil
		}
		return ls.findInMemstoreByID(int64(id))
	} else if err != redis.Nil {
		ls.services.Logger.Println("Error while getting key from memstore: ", err)
	}
	return nil
}

func (ls *Links) createInMemstore(l *Link) {
	if link, err := json.Marshal(l); err == nil {
		pipe := ls.services.IMDS.Client.TxPipeline()
		pipe.Set(context.Background(), "id:"+fmt.Sprint(l.ID), link, 0)
		pipe.Set(context.Background(), "url:"+l.URL, l.ID, 0)
		if _, err := pipe.Exec(context.Background()); err != nil {
			ls.services.Logger.Println("Error while setting link in memstore: ", err)
		}
	} else {
		ls.services.Logger.Println("Error while marshaling link: ", err)
	}
}
