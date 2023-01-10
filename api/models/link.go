package models

import "time"

type Link struct {
	ID        int64     `json:"id"`
	ShortURI  string    `json:"short_uri"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
