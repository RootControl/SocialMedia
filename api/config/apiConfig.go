package apiConfig

import (
	"fmt"
	"net/http"
	"sync/atomic"
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

	return next
}

func (ac *ApiConfig) ApiHitsHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "text/plain; charset=utf-8")
	response.WriteHeader(http.StatusOK)

	response.Write([]byte(fmt.Sprintf("Hits: %d", ac.FileServerHits.Load())))
}

func (ac *ApiConfig) ResetHitsHandler(response http.ResponseWriter, request *http.Request) {
	ac.FileServerHits.Store(0)
	response.Header().Add("Content-Type", "text/plain; charset=utf-8")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Hits reset to 0"))
}
