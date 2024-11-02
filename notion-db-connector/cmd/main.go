package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/taco-tortilla/notion-db-connector/config"
)

func main() {

	config.LoadEnv()

	result := os.Getenv("NOTION_DB_ID")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": result,
		})
	})
	r.Run()
}
