package model

type Link struct {
	LinkID  string `json:"link_id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	UserID  string `json:"user"`
}
