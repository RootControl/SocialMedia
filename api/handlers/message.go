package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RootControl/SocialMedia/api/internal/database"
	"github.com/google/uuid"
)

func (b *BaseHandler) SaveUserMessageHandler(w http.ResponseWriter, r *http.Request) {
	request := struct {
		Body string `json:"body"`
		UserId uuid.UUID `json:"user_id"`
	}{}

	var badRequest BadRequest
	
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		badRequest.Error = err.Error()
		b.sendResponseToClient(w, http.StatusInternalServerError, err.Error(), badRequest)
		return
	}

	if len(request.Body) == 0 {
		badRequest.Error = "message is empty"
		b.sendResponseToClient(w, http.StatusBadRequest, badRequest.Error, badRequest)
		return
	}

	if len(request.Body) > 140 {
		badRequest.Error = "message is to long. Goes over 140 char"
		b.sendResponseToClient(w, http.StatusBadRequest, badRequest.Error, badRequest)
		return
	}

	user, err := b.DbQueries.GetUserById(r.Context(), request.UserId)
	if err != nil {
		badRequest.Error = "user not found"
		b.sendResponseToClient(w, http.StatusBadRequest, badRequest.Error, badRequest)		
		return
	}

	saveMessage := database.SaveMessageParams {
		Body: request.Body,
		UserID: user.ID,
	}
	message, err := b.DbQueries.SaveMessage(r.Context(), saveMessage)
	if err != nil {
		badRequest.Error = "error saving message"
		b.sendResponseToClient(w, http.StatusInternalServerError, err.Error(), badRequest)
		return
	}

	logMessage :=  "message saved for the user: " + user.ID.String()
	b.sendResponseToClient(w, http.StatusCreated, logMessage, message)
}
