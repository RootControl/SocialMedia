package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/RootControl/SocialMedia/api/utils"
)

type Request struct {
	Body string
} 

type Error struct {
	Error string
}

type Response struct {
	Valid bool
}

func ValidateMessageHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	request := Request{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)

	if err != nil {
		errResponse := Error {
			Error: err.Error(),
		}
		log.Print(err.Error())
		utils.RespondWithJSON(w, http.StatusInternalServerError, errResponse)

	} else if len(request.Body) == 0 {
		errResponse := Error {
			Error: "Not able to read the message",
		}
		log.Print(errResponse.Error)
		utils.RespondWithJSON(w, http.StatusBadRequest, errResponse)


	} else if len(request.Body) > 140 {
		errResponse := Error {
			Error: "Message is too long",
		}
		log.Print(errResponse.Error)
		utils.RespondWithJSON(w, http.StatusBadRequest, errResponse)

	} else {
		response := Response {
			Valid: true,
		}
		utils.RespondWithJSON(w, http.StatusOK, response)
	}
}
