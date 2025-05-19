package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RootControl/SocialMedia/api/internal/auth"
	"github.com/RootControl/SocialMedia/api/internal/database"
	"github.com/google/uuid"
)

func (b *BaseHandler) SaveUserMessageHandler(w http.ResponseWriter, r *http.Request) {
	request := struct {
		Body   string    `json:"body"`
		UserId uuid.UUID `json:"user_id"`
	}{}

	var badRequest BadRequest

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		badRequest.Error = err.Error()
		b.sendResponseToClient(w, http.StatusUnauthorized, err.Error(), badRequest)
		return
	}

	userId, err := auth.ValidateJWT(token, b.TokenSecret)
	if err != nil {
		badRequest.Error = err.Error()
		b.sendResponseToClient(w, http.StatusUnauthorized, err.Error(), badRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
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

	user, err := b.DbQueries.GetUserById(r.Context(), userId)
	if err != nil {
		badRequest.Error = "user not found"
		b.sendResponseToClient(w, http.StatusBadRequest, badRequest.Error, badRequest)
		return
	}

	saveMessage := database.SaveMessageParams{
		Body:   request.Body,
		UserID: user.ID,
	}
	message, err := b.DbQueries.SaveMessage(r.Context(), saveMessage)
	if err != nil {
		badRequest.Error = "error saving message"
		b.sendResponseToClient(w, http.StatusInternalServerError, err.Error(), badRequest)
		return
	}

	logMessage := "message saved for the user: " + user.ID.String()
	b.sendResponseToClient(w, http.StatusCreated, logMessage, message)
}

func (b *BaseHandler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	messages, err := b.DbQueries.GetMessages(r.Context())
	if err != nil {
		badRequest := BadRequest{
			Error: "error getting all messages",
		}
		b.sendResponseToClient(w, http.StatusInternalServerError, err.Error(), badRequest)
		return
	}

	b.sendResponseToClient(w, http.StatusOK, "", messages)
}

func (b *BaseHandler) GetMessageByIdHandler(w http.ResponseWriter, r *http.Request) {
	badRequest := BadRequest{}

	if r.PathValue("id") == "" {
		badRequest.Error = "ID is empty"
		b.sendResponseToClient(w, http.StatusBadRequest, badRequest.Error, badRequest)
		return
	}

	messageId, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		badRequest.Error = "Unable to parse ID"
		b.sendResponseToClient(w, http.StatusBadRequest, err.Error(), badRequest)
		return
	}

	message, err := b.DbQueries.GetMessage(r.Context(), messageId)
	if err != nil {
		badRequest.Error = "message not found"
		b.sendResponseToClient(w, http.StatusBadRequest, err.Error(), badRequest)
		return
	}

	b.sendResponseToClient(w, http.StatusOK, "", message)
}
