package main

import (
	"net/http"

	ac "github.com/RootControl/SocialMedia/api/config"
	"github.com/RootControl/SocialMedia/api/handlers"
)

const (
	AppPathRoot = "./app/static"
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
	serverMux.HandleFunc("GET /api/healthz", handlers.AppReadinessHandler)
	serverMux.HandleFunc("POST /api/validate-message", handlers.ValidateMessageHandler)
	serverMux.HandleFunc("POST /api/replace-profane-message", handlers.ReplaceProfaneWordsHandler)

	server := http.Server{
		Addr: ":" + Port,
		Handler: serverMux,
	}

	server.ListenAndServe()
}
