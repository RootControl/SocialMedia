package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/RootControl/SocialMedia/api/utils"
)

var profaneWords = [3]string {
	"kerfuffle",
	"sharbert",
	"fornax",
}

func ReplaceProfaneWordsHandler(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		CleanedBody string `json:"cleaned_body"`
	}
	type Request struct {
		Body string `json:"body"`
	}

	defer r.Body.Close()

	request := Request{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)

	if err != nil {
		log.Printf("error parsing from JSON: %v\n", err.Error())
		utils.RespondWithJSON(w, http.StatusInternalServerError, request)
		return
	}

	response := Response{
		CleanedBody: request.Body,
	}

	for _, badWords := range profaneWords {
		response.CleanedBody = strings.ReplaceAll(response.CleanedBody, badWords, "****")
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}
