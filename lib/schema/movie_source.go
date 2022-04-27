package schema

type MovieSource struct {
	FilmHosterListSelector []string          `json:"film_hoster_list_selector"`
	FilmPostSelector       MoviePostSelector `json:"film_post_selector"`
	FilmListURL            string            `json:"film_list_url"`
	Status                 bool              `json:"status"`
	Name                   string            `json:"name"`
	URL                    string            `json:"url"`
}
