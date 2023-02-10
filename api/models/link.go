package models

import (
	"encoding/json"
	"strings"
	"time"
)

const ValidChars = "bcdfghmnprstvz23456789"

type Link struct {
	id        int64
	ShortURI  string    `json:"short_uri"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

func (l *Link) SetShortURI() {
	l.ShortURI = ShortURIfromID(l.id)
}

func (l *Link) Marshal() ([]byte, error) {
	return json.Marshal(l)
}

func (l *Link) Unmarshal(data []byte) error {
	return json.Unmarshal(data, l)
}

func ShortURIfromID(id int64) string {
	uri := ""
	for id > 0 {
		uri = string(ValidChars[id%int64(len(ValidChars))]) + uri
		id = id / int64(len(ValidChars))
	}
	return uri
}

func IDfromShortURI(uri string) int64 {
	id := int64(0)
	for _, c := range uri {
		id = id*int64(len(ValidChars)) + int64(strings.Index(ValidChars, string(c)))
	}
	return id
}
