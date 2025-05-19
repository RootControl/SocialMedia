package handlers

import (
	"net/http"
)

func (b *BaseHandler) AppReadinessHandler(w http.ResponseWriter, request *http.Request) {
	health := struct{ Status string }{
		Status: "ok",
	}

	b.sendResponseToClient(w, http.StatusOK, "health check ok", health)
}
