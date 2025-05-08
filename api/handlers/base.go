package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/RootControl/SocialMedia/api/internal/database"
)

type BaseHandler struct {
	DbQueries *database.Queries
	
}

type BadRequest struct {
	Error string `json:"error"`
}

func NewBaseHandler(dbConn *sql.DB) *BaseHandler {
	return &BaseHandler{
		DbQueries: database.New(dbConn), 
	}
}

func (b * BaseHandler) sendResponseToClient(w http.ResponseWriter, statusCode int, logMessage string, payload any) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error parsing to JSON: %s\n", err.Error())
	}

	log.Print(logMessage)

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	w.Write(response)
}
