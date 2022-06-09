package schema

type MovieVideo struct {
	SubtitleHosters []string `firestore:"subtitle_hosters" json:"subtitle_hosters"`
	Hosters         []string `firestore:"hosters" json:"hosters"`
	Name            string   `firestore:"name" json:"name"`
}

type MovieArticle struct {
	Description string       `firestore:"description" json:"description"`
	Genders     []string     `firestore:"genders" json:"genders"`
	Videos      []MovieVideo `firestore:"videos" json:"videos"`
	Date        string       `firestore:"date" json:"date"`
	Imdb        string       `firestore:"imdb" json:"imdb"`
}

type MovieArticleSelector struct {
	Description []string             `firestore:"description" json:"description"`
	Genders     []string             `firestore:"genders" json:"genders"`
	Videos      []MovieVideoSelector `firestore:"videos" json:"videos"`
	Date        []string             `firestore:"date" json:"date"`
	Imdb        []string             `firestore:"imdb" json:"imdb"`
}

type MovieVideoSelector struct {
	Hosters []string `firestore:"hosters" json:"hosters"`
}
