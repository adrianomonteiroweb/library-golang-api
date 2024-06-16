package models

type Author struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DateBirth   string `json:"date_birth"`
	Nationality string `json:"nationality"`
}