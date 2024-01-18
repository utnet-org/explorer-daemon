package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// func to get env value
func EnvLoad(key string) string {
	// load .env file
	err := godotenv.Load(EnvFile)
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
