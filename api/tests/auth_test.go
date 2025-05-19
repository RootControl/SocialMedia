package tests

import (
	"net/http"
	"testing"
	"time"

	"github.com/RootControl/SocialMedia/api/internal/auth"
	"github.com/google/uuid"
)

func TestAuthJWT(t *testing.T) {
	userId := uuid.New()
	tokenSecrect := "test"
	expireIn := 1 * time.Hour

	token, err := auth.MakeJWT(userId, tokenSecrect, expireIn)
	if err != nil {
		t.Error(err)
		return
	}

	id, err := auth.ValidateJWT(token, tokenSecrect)
	if err != nil {
		t.Error(err)
		return
	}

	if userId != id {
		t.Errorf("IDs not match | userId: %s - Id: %s", userId, id)
		return
	}
}

func TestBearerToken(t *testing.T) {
	headers := make(http.Header)
	headers.Add("Authorization", "Bearer token")

	_, err := auth.GetBearerToken(headers)
	if err != nil {
		t.Errorf("Unable to get the token: %s", err.Error())
		return
	}
}
