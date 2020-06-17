package main

import (
	"charactersManager/controller"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":3000"
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello World")
	})
	router.HandleFunc("/characters", controller.GetCharacters).Methods("GET")

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
