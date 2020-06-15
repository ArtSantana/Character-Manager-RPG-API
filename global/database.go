package global

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Just for remove the warning
)

func init() {
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s)/%s", Env("USER"), Env("PASSWORD"), Env("IP"), Env("DATABASE"))

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		fmt.Println(err)
		log.Fatalf("Error connecting to the database")
	}

	defer db.Close()

	fmt.Println("Succesfully connected to the database")
}
