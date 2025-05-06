package main

import (
	"encoding/json"
	"log"
	"net/http"

	ac "github.com/RootControl/SocialMedia/api/config"
)

const (
	AppPathRoot = "./app/static"
	AdminPathRoot = "./app/admin"
	Port = "8080"
)

func main() {
	serverMux := http.NewServeMux()
	apiConfig := ac.NewApiConfig()

	// Front-end handlers
	staticFilesHandler := http.StripPrefix("/app", http.FileServer(http.Dir(AppPathRoot)))
	serverMux.Handle("/app/", apiConfig.MiddlewareMetricsIncrement(staticFilesHandler))

	// Internal administrative use
	serverMux.HandleFunc("GET /admin/metrics", apiConfig.ApiHitsHandler)
	serverMux.HandleFunc("POST /admin/reset", apiConfig.ResetHitsHandler)

	// Back-end handlers
	serverMux.HandleFunc("GET /api/healthz", appReadinessHandler)
	serverMux.HandleFunc("POST /api/validate-message", validateMessageHandler)

	server := http.Server{
		Addr: ":" + Port,
		Handler: serverMux,
	}

	server.ListenAndServe()
}

func appReadinessHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "text/plain; charset=utf-8")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(http.StatusText(http.StatusOK)))
}

func validateMessageHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	request := struct{ Body string }{}
	errorResponse := struct{ Error string }{}
	response := struct{ Valid bool }{ Valid: true }

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)

	if err != nil {
		errorResponse.Error = err.Error()
		log.Print(err.Error())
		respondWithJSON(w, http.StatusInternalServerError, errorResponse)

	} else if len(request.Body) == 0 {
		errorResponse.Error = "Not able to read the message"
		log.Print(errorResponse.Error)
		respondWithJSON(w, http.StatusBadRequest, errorResponse)


	} else if len(request.Body) > 140 {
		errorResponse.Error = "Message is too long"
		log.Print(errorResponse.Error)
		respondWithJSON(w, http.StatusBadRequest, errorResponse)

	} else {
		respondWithJSON(w, http.StatusOK, response)
	}
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload any) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	w.Write(response)
	return nil
}
