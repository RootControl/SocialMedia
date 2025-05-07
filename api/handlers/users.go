package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/RootControl/SocialMedia/api/utils"
)

func (b *BaseHandler) CreateNewUserHandler(w http.ResponseWriter, r *http.Request) {
	request := struct { Email string `json:"email"`}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		log.Printf("error decoding JSON: %s", err.Error())
		utils.RespondWithJSON(w, http.StatusInternalServerError, request)
	}

	user, err := b.DbQueries.CreateUser(r.Context(), request.Email)
	if err != nil {
		log.Printf("error creating user: %s", err.Error())
		utils.RespondWithJSON(w, http.StatusInternalServerError, request)
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}
