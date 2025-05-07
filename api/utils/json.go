package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload any) error {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error parsing to JSON: %s\n", err.Error())
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	w.Write(response)

	return nil
}
