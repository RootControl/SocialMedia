package main

import (
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
