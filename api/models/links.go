package models

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/h00s-go/tiny-link-backend/api/services"
	"github.com/h00s-go/tiny-link-backend/db/sql"
)

type Link struct {
	ID        int64     `json:"id"`
	ShortURI  string    `json:"short_uri"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

func GetLinkByID(s *services.Services, id string) (*Link, error) {
	l := &Link{}

	if err := s.DB.Conn.QueryRow(context.Background(), sql.GetLinkByID, id).Scan(&l.ID, &l.ShortURI, &l.URL, &l.CreatedAt); err != nil {
		return nil, err
	}

	return l, nil
}

func GetLinkByShortURI(s *services.Services, shortURI string) (*Link, error) {
	l := &Link{}

	value, err := s.IMDS.Client.Get(context.Background(), shortURI).Result()
	if err != nil && err != redis.Nil {
		s.Logger.Println("Error while getting key from memstore: ", err)
	} else {
		l.URL = value
		return l, nil
	}

	if err := s.DB.Conn.QueryRow(context.Background(), sql.GetLinkByShortURI, shortURI).Scan(&l.ID, &l.ShortURI, &l.URL, &l.CreatedAt); err != nil {
		return nil, err
	}

	if err := s.IMDS.Client.Set(context.Background(), shortURI, l.URL, 0).Err(); err != nil {
		s.Logger.Println("Error while setting key to memstore: ", err)
	}

	return l, nil
}
