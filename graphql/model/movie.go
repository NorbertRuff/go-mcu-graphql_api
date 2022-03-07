package model

type Movie struct {
	MovieID          string `json:"movie_id"`
	Title            string `json:"title"`
	ReleaseDate      string `json:"release_date"`
	BoxOffice        string `json:"box_office"`
	Duration         int    `json:"duration"`
	Overview         string `json:"overview"`
	CoverURL         string `json:"cover_url"`
	TrailerURL       string `json:"trailer_url"`
	DirectedBy       string `json:"directed_by"`
	Phase            int    `json:"phase"`
	Saga             string `json:"saga"`
	Chronology       int    `json:"chronology"`
	PostCreditScenes string `json:"post_credit_scenes"`
	ImdbID           string `json:"imdb_id"`
	UserID           string `json:"user"`
}
