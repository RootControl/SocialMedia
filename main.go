package main

import "net/http"

const (
	FilePathRoot = "./internal/static/"
	Port = "8080"
)

func main() {
	serverMux := http.NewServeMux()

	staticFilesHandler := http.StripPrefix("/app", http.FileServer(http.Dir(FilePathRoot)))

	serverMux.Handle("/app/", staticFilesHandler)
	serverMux.HandleFunc("/healthz", appReadinessHandler)

	server := http.Server{
		Addr: ":" + Port,
		Handler: serverMux,
	}

	server.ListenAndServe()
}

func appReadinessHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "text/plain charset=utf-8")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(http.StatusText(http.StatusOK)))
}
