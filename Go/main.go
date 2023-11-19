package main

import (
	"github.com/joho/godotenv"
	"jdreese/setlistplay/api/router"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router.Init()
	router.Run()
}
