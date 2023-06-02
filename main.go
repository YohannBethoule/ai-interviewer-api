package main

import (
	"github.com/joho/godotenv"
	"interviewer-api/router"
)

func main() {
	godotenv.Load()

	router.InitRouter()
}
