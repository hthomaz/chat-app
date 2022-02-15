package main

import (
	"log"
	"os"

	"heitor/chatApp"

	"github.com/joho/godotenv"
)

func main() {

	//Load enviroments data
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	redisURL := os.Getenv("REDIS_URL")
	port := os.Getenv("PORT")
	chatApp.Obverser(redisURL, port)
}
