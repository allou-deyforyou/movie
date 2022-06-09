package schema

type MovieCategory string

const (
	MovieFilm  MovieCategory = "film"
	MovieSerie MovieCategory = "serie"
	MovieManga MovieCategory = "manga"
)

type MoviePost struct {
	Category MovieCategory `json:"category"`
	Source   string        `json:"source"`
	Title    string        `json:"title"`
	Image    string        `json:"image"`
	Link     string        `json:"link"`
}

type MoviePostSelector struct {
	Title []string `firestore:"title" json:"title"`
	Image []string `firestore:"image" json:"image"`
	List  []string `firestore:"list" json:"list"`
	Link  []string `firestore:"link" json:"link"`
}
