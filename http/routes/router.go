package routes

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"parser/http/controllers"
	"time"
)

type Route struct {
	Router *mux.Router
	Serv   *http.Server
	Action controllers.App
}

func (a *Route) CreateRoute() {
	a.initializeRoutes()
}
func (a *Route) initializeRoutes() {
	a.Router.HandleFunc("/users", a.Action.GetUsers).Methods("GET")
	a.Router.HandleFunc("/user/register", a.Action.UserRegister).Methods("POST")
	a.Router.HandleFunc("/user/{id:[0-9]+}", a.Action.GetUser).Methods("GET")
	a.Router.HandleFunc("/user/update/{id:[0-9]+}", a.Action.UpdateUser).Methods("PUT")
	a.Router.HandleFunc("/user/delete/{id:[0-9]+}", a.Action.DeleteUser).Methods("DELETE")
}

func (a *Route) Run(addr string) {
	a.Serv = &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      a.Router,
	}
	log.Fatal(a.Serv.ListenAndServe())
}
