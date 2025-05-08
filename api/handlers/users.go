package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (b *BaseHandler) CreateNewUserHandler(w http.ResponseWriter, r *http.Request) {
	request := struct { Email string `json:"email"`}{}
	var badRequest BadRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		badRequest.Error = "error parsing JSON"
		b.sendResponseToClient(w, http.StatusInternalServerError, err.Error(), badRequest)
	}

	user, err := b.DbQueries.CreateUser(r.Context(), request.Email)
	if err != nil {
		badRequest.Error = "error creating user"
		b.sendResponseToClient(w, http.StatusInternalServerError, err.Error(), badRequest)
	}

	logMessage := fmt.Sprintf("user %v created", user.ID)
	b.sendResponseToClient(w, http.StatusCreated, logMessage, user)
}
