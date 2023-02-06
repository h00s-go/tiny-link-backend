package models

import (
	"encoding/json"
	"time"
)

type Link struct {
	ID             int64     `json:"id"`
	ShortURI       string    `json:"short_uri"`
	URL            string    `json:"url"`
	CreatedAt      time.Time `json:"created_at"`
	LastAccessedAt time.Time `json:"last_accessed_at"`
	AccessCount    int64     `json:"access_count"`
}

func (l *Link) GenerateShortURI() {
	const validChars = "ABCDEFHJKLMNPRSTUVXYZabcdefgijkmnprstuvxyz23456789"
	uri := ""
	id := l.ID
	for id > 0 {
		uri = string(validChars[id%int64(len(validChars))]) + uri
		id = id / int64(len(validChars))
	}
	l.ShortURI = uri
}

func (l *Link) Marshal() ([]byte, error) {
	return json.Marshal(l)
}

func (l *Link) Unmarshal(data []byte) error {
	return json.Unmarshal(data, l)
}
