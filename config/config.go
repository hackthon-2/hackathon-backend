package config

import (
	"github.com/joho/godotenv"
	"log"
)

// Config func to get env value
func Config(){
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}