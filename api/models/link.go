package models

import (
	"encoding/json"
	"strings"
	"time"
)

const ValidChars = "bcdfghmnprstvz23456789"

type Link struct {
	ID        int64     `json:"-"`
	ShortURI  string    `json:"short_uri"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type AliasLink Link
type PublicLink struct {
	ShortURI string `json:"short_uri"`
	*AliasLink
}

func (link *Link) MarshalJSON() ([]byte, error) {
	return json.Marshal(&PublicLink{
		ShortURI:  ShortURIfromID(link.ID),
		AliasLink: (*AliasLink)(link),
	})
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
