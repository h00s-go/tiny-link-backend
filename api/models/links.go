package models

import (
	"context"

	"github.com/go-redis/redis/v9"
	"github.com/h00s-go/tiny-link-backend/api/services"
	"github.com/h00s-go/tiny-link-backend/db/sql"
)

type Links struct {
	services *services.Services
}

func NewLinks(services *services.Services) *Links {
	return &Links{
		services: services,
	}
}

func (ls *Links) FindByID(id string) (*Link, error) {
	l := &Link{}

	if err := ls.services.DB.Conn.QueryRow(context.Background(), sql.GetLinkByID, id).Scan(&l.ID, &l.ShortURI, &l.URL, &l.CreatedAt); err != nil {
		return nil, err
	}

	return l, nil
}

func (ls *Links) FindByShortURI(shortURI string) (*Link, error) {
	if l := ls.findInMemstoreByShortURI(shortURI); l != nil {
		return l, nil
	}

	l := &Link{}
	if err := ls.services.DB.Conn.QueryRow(context.Background(), sql.GetLinkByShortURI, shortURI).Scan(&l.ID, &l.ShortURI, &l.URL, &l.CreatedAt); err != nil {
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
	if err := ls.services.DB.Conn.QueryRow(context.Background(), sql.GetLinkByURL, URL).Scan(&l.ID, &l.ShortURI, &l.URL, &l.CreatedAt); err != nil {
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
	l.GenerateShortName()

	if _, err := tx.Exec(context.Background(), sql.UpdateLinkShortURI, l.ShortURI, l.ID); err != nil {
		tx.Rollback(context.Background())
		return nil, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	return l, nil
}

// ++++++ Memstore ++++++

func (ls *Links) findInMemstoreByShortURI(shortURI string) *Link {
	l := &Link{}

	value, err := ls.services.IMDS.Client.Get(context.Background(), "short_uri:"+shortURI).Result()
	if err == nil {
		if l.Unmarshal([]byte(value)) != nil {
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
	shortURI, err := ls.services.IMDS.Client.Get(context.Background(), "url:"+url).Result()
	if err == nil {
		return ls.findInMemstoreByShortURI(shortURI)
	} else if err != redis.Nil {
		ls.services.Logger.Println("Error while getting key from memstore: ", err)
	}
	return nil
}

func (ls *Links) createInMemstore(l *Link) {
	if link, err := l.Marshal(); err == nil {
		pipe := ls.services.IMDS.Client.TxPipeline()
		pipe.Set(context.Background(), "short_uri:"+l.ShortURI, link, 0)
		pipe.Set(context.Background(), "url:"+string(l.URL), l.ShortURI, 0)
		if _, err := pipe.Exec(context.Background()); err != nil {
			ls.services.Logger.Println("Error while setting link in memstore: ", err)
		}
	} else {
		ls.services.Logger.Println("Error while marshaling link: ", err)
	}
}
