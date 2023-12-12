package helpers

import "os"

var BasePath string

func init() {
	BasePath = os.Getenv("BASE_PATH")
	if BasePath == "" {
		BasePath = "/"
	}
}
