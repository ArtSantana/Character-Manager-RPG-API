package global

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Just for remove the warning
)

var connectionString string = fmt.Sprintf("%s:%s@tcp(%s)/%s", Env("USER"), Env("PASSWORD"), Env("IP"), Env("DATABASE"))
var db, err = sql.Open("mysql", connectionString)

type Character struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level string `json:"level"`
	Class string `json:"class"`
	Race  string `json:"race"`
}

func init() {

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Succesfully connected to the database")
}

func SelectQuery(table string, condition string) []Character {
	var (
		characters []Character
	)

	var queryString string

	if condition != "" {
		queryString = fmt.Sprintf("SELECT * FROM %s WHERE %s", table, condition)
	} else {
		queryString = fmt.Sprintf("SELECT * FROM %s", table)
	}
	result, err := db.Query(queryString)

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var character Character

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
