package repository

import (
	"charactersManager/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SelectQuery(table string, condition string) []models.Character {
	var (
		characters []models.Character
	)

	var queryString string

	if condition != "" {
		queryString = fmt.Sprintf("SELECT * FROM %s WHERE %s;", table, condition)
	} else {
		queryString = fmt.Sprintf("SELECT * FROM %s;", table)
	}
	result, err := db.Query(queryString)

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var character models.Character

		err = result.Scan(
			&character.ID,
			&character.Name,
			&character.Level,
			&character.Class,
			&character.Race,
		)

		if err != nil {
			panic(err.Error())
		}

		characters = append(characters, character)

	}
	return characters
}
