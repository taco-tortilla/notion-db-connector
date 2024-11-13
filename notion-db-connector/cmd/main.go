package main

import (
	"github.com/gin-gonic/gin"
	"github.com/taco-tortilla/notion-db-connector/config"
	"github.com/taco-tortilla/notion-db-connector/internal/infrastructure/api"
	"github.com/taco-tortilla/notion-db-connector/internal/interfaces/handler"
	"github.com/taco-tortilla/notion-db-connector/internal/usecase"
)

func main() {
	config.LoadEnv()

	feedbackAPI := api.NewFeedbackAPI()
	feedbackUseCase := usecase.NewFeedbackUseCase(feedbackAPI)
	feedbackHandler := handler.NewFeedbackHandler(feedbackUseCase)

	r := gin.Default()
	r.POST("/feedback", feedbackHandler.AddFeedback)
	r.Run()
}
