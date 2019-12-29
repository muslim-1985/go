package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"os"
	"parser/config"
	"parser/users/controllers"
	"parser/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a := config.App{}
	//_ = config.InitWorkers
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	userModule := &controllers.App{DB: a.DB}
	route := mux.NewRouter()
	r := routes.Route{UserAction: *userModule, Router: route}
	r.CreateRoute()
	//quoteChan := parser.Grab()
	//for i := 0; i < 5; i++ { //получаем 5 цитат и закругляемся
	//	fmt.Println(<-quoteChan, "\n")
	//}
	r.Run("localhost:8080")

}
