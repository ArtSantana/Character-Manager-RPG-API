package models

type Character struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level string `json:"level"`
	Class string `json:"class"`
	Race  string `json:"race"`
}
