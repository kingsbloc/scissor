package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(variable string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	result := os.Getenv(variable)
	if result == "" {
		log.Fatal("You must set your " + variable + " environmental variable")
	}
	return result
}
