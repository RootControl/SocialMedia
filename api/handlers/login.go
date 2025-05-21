package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/RootControl/SocialMedia/api/internal/auth"
	"github.com/RootControl/SocialMedia/api/internal/database"
	"github.com/google/uuid"
)

func (b *BaseHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	request := struct {
		Email        string `json:"email"`
		Hashpassword string `json:"password"`
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

	token, err := auth.MakeJWT(user.ID, b.TokenSecret, time.Duration(time.Hour))
	if err != nil {
		badRequest.Error = "unable to validate user"
		b.sendResponseToClient(w, http.StatusUnauthorized, err.Error(), badRequest)
		return
	}

	refresh := auth.MakeRefreshToken()

	refreshToken := database.SaveRefreshTokenParams{
		Token:     refresh,
		UserID:    user.ID,
		ExpiredAt: time.Now().AddDate(0, 0, 60).UTC(),
	}

	rToken, err := b.DbQueries.SaveRefreshToken(r.Context(), refreshToken)
	if err != nil {
		badRequest.Error = "unable to validate user"
		b.sendResponseToClient(w, http.StatusUnauthorized, err.Error(), badRequest)
		return
	}

	response := struct {
		ID           uuid.UUID `json:"id"`
		Email        string    `json:"email"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		Token        string    `json:"token"`
		RefreshToken string    `json:"refresh_token"`
	}{
		ID:           user.ID,
		Email:        user.Email,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Token:        token,
		RefreshToken: rToken.Token,
	}

	b.sendResponseToClient(w, http.StatusOK, "logged in", response)
}
