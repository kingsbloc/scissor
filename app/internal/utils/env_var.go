package utils

import (
	"log"
	"os"
)

func GetEnvVar(variable string) string {
	result := os.Getenv(variable)
	if result == "" {
		log.Fatal("You must set your " + variable + " environmental variable")
	}
	return result
}
