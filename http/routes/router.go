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
	a.Router.HandleFunc("/products", a.Action.GetProducts).Methods("GET")
	a.Router.HandleFunc("/product", a.Action.CreateProduct).Methods("POST")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.Action.GetProduct).Methods("GET")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.Action.UpdateProduct).Methods("PUT")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.Action.DeleteProduct).Methods("DELETE")
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
