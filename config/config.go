package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load("dev.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
