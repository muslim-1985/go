package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"os"
	"parser/http/controllers"
	"parser/http/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a := controllers.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	route := mux.NewRouter()
	r := routes.Route{Action: a, Router: route}
	r.CreateRoute()
	r.Run("localhost:8080")
}
