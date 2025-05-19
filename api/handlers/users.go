package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/RootControl/SocialMedia/api/internal/auth"
	"github.com/RootControl/SocialMedia/api/internal/database"
	"net/http"
)

func (b *BaseHandler) CreateNewUserHandler(w http.ResponseWriter, r *http.Request) {
	request := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	var badRequest BadRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		badRequest.Error = "error parsing JSON"
		b.sendResponseToClient(w, http.StatusInternalServerError, err.Error(), badRequest)
		return
	}

	dbUser, err := b.DbQueries.GetUserByEmail(r.Context(), request.Email)
	if err == nil && dbUser.Email != "" {
		badRequest.Error = "User already exists"
		b.sendResponseToClient(w, http.StatusInternalServerError, badRequest.Error, badRequest)
		return
	}

	hashPassword, err := auth.HashPassword(request.Password)
	if err != nil {
		badRequest.Error = "unable to hash the password"
		b.sendResponseToClient(w, http.StatusInternalServerError, err.Error(), badRequest)
		return
	}

	userParam := database.CreateUserParams{
		Email:        request.Email,
		HashPassword: hashPassword,
	}
	user, err := b.DbQueries.CreateUser(r.Context(), userParam)
	if err != nil {
		badRequest.Error = "error creating user"
		b.sendResponseToClient(w, http.StatusInternalServerError, err.Error(), badRequest)
		return
	}

	response := database.GetUserByIdRow{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	logMessage := fmt.Sprintf("user %v created", user.ID)
	b.sendResponseToClient(w, http.StatusCreated, logMessage, response)
}
