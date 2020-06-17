package controller

import (
	"charactersManager/repository"
	"encoding/json"
	"net/http"
)

func GetCharacters(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	characters := repository.SelectQuery("characters", "")

	result, err := json.Marshal(characters)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling the characters"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(result)
}
