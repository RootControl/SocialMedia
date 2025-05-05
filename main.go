package main

import (
	"net/http"

	ac "github.com/RootControl/SocialMedia/api/config"
)

const (
	FilePathRoot = "./app/static"
	Port = "8080"
)

func main() {
	serverMux := http.NewServeMux()
	apiConfig := ac.NewApiConfig()

	staticFilesHandler := http.StripPrefix("/app", http.FileServer(http.Dir(FilePathRoot)))

	serverMux.Handle("/app/", apiConfig.MiddlewareMetricsIncrement(staticFilesHandler))

	serverMux.HandleFunc("GET /api/healthz", appReadinessHandler)
	serverMux.HandleFunc("GET /api/metrics", apiConfig.ApiHitsHandler)
	serverMux.HandleFunc("GET /api/reset", apiConfig.ResetHitsHandler)

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
