package hoster

import (
	"net/http"
	"server/lib/schema"
)

type Hoster interface {
	SetData(schema.MovieHoster, *http.Client)
	Video(string, string) *schema.MovieVideo
	Verify(string) bool
}

func GetAllHosters() []Hoster {
	return []Hoster{&UserloadHoster{}, &UqloadHoster{}}
}
