package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("failed to load .env: %v", err)
	}
}
