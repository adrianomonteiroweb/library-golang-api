package models

type Book struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	AuthorID     int    `json:"author_id"`
	PublisheYear int    `json:"publishe_year"`
	Genre        string `json:"genre"`
}