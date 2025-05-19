package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/RootControl/SocialMedia/api/internal/auth"
	"github.com/google/uuid"
)

func (b *BaseHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	request := struct {
		Email            string `json:"email"`
		Hashpassword     string `json:"password"`
		ExpiresInSeconds int    `json:"expires_in_seconds"`
	}{}

	badRequest := BadRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		badRequest.Error = "Unable to decode the JSON request"
		b.sendResponseToClient(w, http.StatusBadRequest, err.Error(), badRequest)
		return
	}

	user, err := b.DbQueries.GetUserByEmail(r.Context(), request.Email)
	if err != nil {
		badRequest.Error = "Incorrect email or password"
		b.sendResponseToClient(w, http.StatusUnauthorized, err.Error(), badRequest)
		return
	}

	err = auth.CheckPasswordHash(user.HashPassword, request.Hashpassword)
	if err != nil {
		badRequest.Error = "Incorrect email or password"
		b.sendResponseToClient(w, http.StatusUnauthorized, err.Error(), badRequest)
		return
	}

	if request.ExpiresInSeconds == 0 || request.ExpiresInSeconds > 3600 {
		request.ExpiresInSeconds = int(time.Hour)
	}

	token, err := auth.MakeJWT(user.ID, b.TokenSecret, time.Duration(request.ExpiresInSeconds))
	if err != nil {
		badRequest.Error = "unable to validate user"
		b.sendResponseToClient(w, http.StatusUnauthorized, err.Error(), badRequest)
		return
	}

	response := struct {
		ID        uuid.UUID `json:"id"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Token     string    `json:"token"`
	}{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Token:     token,
	}

	b.sendResponseToClient(w, http.StatusOK, "logged in", response)
}
