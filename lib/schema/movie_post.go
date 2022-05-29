package schema

type MoviePost struct {
	Source string `json:"source"`
	Title  string `json:"title"`
	Image  string `json:"image"`
	Link   string `json:"link"`
}

type MoviePostSelector struct {
	Title []string `json:"title"`
	Image []string `json:"image"`
	List  []string `json:"list"`
	Link  []string `json:"link"`
}
