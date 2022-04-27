package schema

const (
	MOVIE_SOURCES_COLLECTION = "moviesources"
)

type MoviePost struct {
	Source string `json:"source"`
	Title  string `json:"title"`
	Image  string `json:"image"`
	Link   string `json:"link"`
}

type MoviePostSelector struct {
	Title string   `json:"title"`
	Image []string `json:"image"`
	List  string   `json:"list"`
	Link  []string `json:"link"`
}

type MoviePostListResponse struct {
	Data []MoviePost `json:"data"`
}

type MoviePostListReqest struct {
	Query string `json:"query"`
	Page  int    `json:"page"`
}
