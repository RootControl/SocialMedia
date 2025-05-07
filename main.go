package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	ac "github.com/RootControl/SocialMedia/api/config"
	"github.com/RootControl/SocialMedia/api/handlers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	AppPathRoot = "./app/static"
	Port = "8080"
)

func main() {
	godotenv.Load()
	dbUrl := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Printf("error loading database: %s", err.Error())
		return 
	}
	defer db.Close()

	serverMux := http.NewServeMux()
	apiConfig := ac.NewApiConfig()
	baseHandler := handlers.NewBaseHandler(db)

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
	serverMux.HandleFunc("POST /api/users", baseHandler.CreateNewUserHandler)

	server := http.Server{
		Addr: ":" + Port,
		Handler: serverMux,
	}

	server.ListenAndServe()
}
