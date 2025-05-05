package apiConfig

import (
	"net/http"
	"sync/atomic"
)

type ApiConfig struct {
	FileServerHits atomic.Uint64
}

func (ac *ApiConfig) MiddlewareMetricsIncrement(next http.Handler) http.Handler {
	_ = ac.FileServerHits.Add(1)

	return next
}
