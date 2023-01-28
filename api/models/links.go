package models

import (
	"context"
	"encoding/json"

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
	l := &Link{}

	value, err := ls.services.IMDS.Client.Get(context.Background(), "short_uri:"+shortURI).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(value), &l); err != nil {
			ls.services.Logger.Println("Error while unmarshaling link: ", err)
		}
		return l, nil
	} else if err != redis.Nil {
		ls.services.Logger.Println("Error while getting key from memstore: ", err)
	}

	if err := ls.services.DB.Conn.QueryRow(context.Background(), sql.GetLinkByShortURI, shortURI).Scan(&l.ID, &l.ShortURI, &l.URL, &l.CreatedAt); err != nil {
		return nil, err
	}

	link, err := json.Marshal(l)
	if err == nil {
		if err := ls.services.IMDS.Client.Set(context.Background(), "short_uri:"+shortURI, link, 0).Err(); err != nil {
			ls.services.Logger.Println("Error while setting key to memstore: ", err)
		}
		if err := ls.services.IMDS.Client.Set(context.Background(), "url:"+string(l.URL), link, 0).Err(); err != nil {
			ls.services.Logger.Println("Error while setting key to memstore: ", err)
		}
	} else {
		ls.services.Logger.Println("Error while marshaling link: ", err)
	}

	return l, nil
}

func (ls *Links) Create(l *Link) error {
	shortURI, err := ls.services.IMDS.Client.Get(context.Background(), "url:"+l.URL).Result()
	if err == nil {
		link, err := ls.services.IMDS.Client.Get(context.Background(), "short_uri:"+shortURI).Result()
		if err == nil {
			if err := json.Unmarshal([]byte(link), &l); err != nil {
				ls.services.Logger.Println("Error while unmarshaling link: ", err)
			}
			return nil
		}
	} else if err != redis.Nil {
		ls.services.Logger.Println("Error while getting key from memstore: ", err)
	}

	h := NewHost(l.URL)
	if err := h.IsValid(); err != nil {
		return err
	}

	tx, err := ls.services.DB.Conn.Begin(context.Background())
	if err != nil {
		return err
	}

	if err := tx.QueryRow(context.Background(), sql.CreateLink, l.URL).Scan(&l.ID); err != nil {
		tx.Rollback(context.Background())
		return err
	}
	l.GenerateShortName()

	if _, err := tx.Exec(context.Background(), sql.UpdateLinkShortURI, l.ShortURI, l.ID); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return err
	}

	link, err := json.Marshal(l)
	if err == nil {
		if err := ls.services.IMDS.Client.Set(context.Background(), "shorturi:"+l.ShortURI, link, 0).Err(); err != nil {
			ls.services.Logger.Println("Error while setting key to memstore: ", err)
		}
		if err := ls.services.IMDS.Client.Set(context.Background(), "url:"+string(l.URL), l.ShortURI, 0).Err(); err != nil {
			ls.services.Logger.Println("Error while setting key to memstore: ", err)
		}
	} else {
		ls.services.Logger.Println("Error while marshaling link: ", err)
	}

	return nil
}
