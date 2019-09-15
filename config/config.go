package config

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	DB     *sql.DB
	Router *mux.Router
	Serv   *http.Server
}
