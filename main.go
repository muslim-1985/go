package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"os"
	"parser/config"
	"parser/http/controllers"
	"parser/http/parsing"
	"parser/http/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a := controllers.InitDB
	_ = config.InitWorkers
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	route := mux.NewRouter()
	r := routes.Route{Action: a, Router: route}
	r.CreateRoute()
	quoteChan := parser.Grab()
	for i := 0; i < 5; i++ { //получаем 5 цитат и закругляемся
		fmt.Println(<-quoteChan, "\n")
	}
	r.Run("localhost:8080")

}
