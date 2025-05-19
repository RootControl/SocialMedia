package apiConfig

import (
	"net/http"
	"sync/atomic"
	"text/template"
)

type ApiConfig struct {
	FileServerHits atomic.Uint32
}

func NewApiConfig() *ApiConfig {
	return &ApiConfig{}
}

func (ac *ApiConfig) MiddlewareMetricsIncrement(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ac.FileServerHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (ac *ApiConfig) ApiHitsHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "text/html; charset=utf-8")

	template, err := template.ParseFiles("./app/admin/metrics/hits.html")
	if err != nil {
		http.Error(response, "Error loading template", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	data := struct{ Hits uint32 }{Hits: ac.FileServerHits.Load()}

	template.Execute(response, data)
}

func (ac *ApiConfig) ResetHitsHandler(response http.ResponseWriter, request *http.Request) {
	ac.FileServerHits.Store(0)
	response.Header().Add("Content-Type", "text/plain; charset=utf-8")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Hits reset to 0"))
}
