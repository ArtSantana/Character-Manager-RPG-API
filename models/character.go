package models

type Character struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
	Class string `json:"class"`
	Race  string `json:"race"`
}

type CharacterNoID struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
	Class string `json:"class"`
	Race  string `json:"race"`
}
