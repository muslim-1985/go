package main
import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"bot/http/controllers"
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
	    a.Run("localhost:8080")
}
