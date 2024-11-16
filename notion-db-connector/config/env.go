package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	DbId          string
	NotionURL     string
	NotionVersion string
	NotionToken   string
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("failed to load .env: %v", err)
	}

	DbId = os.Getenv("NOTION_DB_ID")
	NotionURL = os.Getenv("NOTION_BASE_URL")
	NotionVersion = os.Getenv("NOTION_VERSION")
	NotionToken = os.Getenv("NOTION_TOKEN")
}
