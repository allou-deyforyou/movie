package schema


const (
	MOVIE_HOSTERS_COLLECTION = "moviehosters"
)


type MovieHoster struct {
	Status bool   `json:"status"`
	Name   string `json:"name"`
	URL    string `json:"url"`
}

type MovieVideo struct {
	Referer string `json:"referer"`
	Link    string `json:"link"`
}
