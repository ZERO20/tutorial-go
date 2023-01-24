package models

// estructura
// TAGS `json:"id"`
type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Phone  int    `json:"phone"`
	Status bool   `json:"status"`
}
