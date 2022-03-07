package model

type TvShow struct {
	ShowID         string `json:"show_id"`
	Title          string `json:"title"`
	ReleaseDate    string `json:"release_date"`
	LastAiredDate  string `json:"last_aired_date"`
	NumberSeasons  int    `json:"number_seasons"`
	NumberEpisodes int    `json:"number_episodes"`
	Overview       string `json:"overview"`
	CoverURL       string `json:"cover_url"`
	TrailerURL     string `json:"trailer_url"`
	DirectedBy     string `json:"directed_by"`
	Phase          int    `json:"phase"`
	Saga           string `json:"saga"`
	ImdbID         string `json:"imdb_id"`
	UserID         string `json:"user"`
}
