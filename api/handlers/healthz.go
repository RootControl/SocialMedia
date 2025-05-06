package handlers

import (
	"log"
	"net/http"

	"github.com/RootControl/SocialMedia/api/utils"
)

type HealthStatus struct {
	Status string
}

func AppReadinessHandler(w http.ResponseWriter, request *http.Request) {
	health := HealthStatus {
		Status: "ok",
	}

	err := utils.RespondWithJSON(w, http.StatusOK, health)
	if err != nil {
		log.Printf("error creating JSON: %v", err.Error())
	}
}
