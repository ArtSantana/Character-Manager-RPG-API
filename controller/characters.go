package controller

import (
	"charactersManager/models"
	"charactersManager/repository"
	"encoding/json"
	"fmt"
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

func PostCharacter(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(req.Body)

	var character models.CharacterNoID

	err := decoder.Decode(&character)

	if err != nil {
		panic(err)
	}

	valid, error := isValid(character)

	resBody := fmt.Sprintf(`{"error": "%s"}`, error)

	if !valid {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(resBody))

		return
	}

	isError := repository.PostCharacterInDatabase(character)

	if isError {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusCreated)
}

func isValid(character models.CharacterNoID) (bool, string) {
	if character.Name == "" {
		return false, "O campo name é obrigatório!"
	}

	if character.Level < 1 {
		return false, "O campo level é obrigatório!"
	}

	if character.Class == "" {
		return false, "O campo class é obrigatório!"
	}

	if character.Race == "" {
		return false, "O campo race é obrigatório!"
	}

	return true, ""
}
