package helpers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var BasePath string

func init() {
	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatal("Can not load env variables")
	}

	BasePath = os.Getenv("BASE_PATH")
	if BasePath == "" {
		BasePath = "/"
	}
}
