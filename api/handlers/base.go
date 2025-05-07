package handlers

import (
	"database/sql"

	"github.com/RootControl/SocialMedia/api/internal/database"
)

type BaseHandler struct {
	DbQueries *database.Queries
	
}

func NewBaseHandler(dbConn *sql.DB) *BaseHandler {
	return &BaseHandler{
		DbQueries: database.New(dbConn), 
	}
}
