package handlers

import (
	"net/http"

	"github.com/RootControl/SocialMedia/api/internal/auth"
)

func (b *BaseHandler) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	badRequest := BadRequest{}

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

	refreshToken, err := b.DbQueries.GetRefreshToken(r.Context(), userId)
	if err != nil {

	}
}
