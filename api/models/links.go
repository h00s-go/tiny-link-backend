package models

import (
	"context"
	"time"

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
