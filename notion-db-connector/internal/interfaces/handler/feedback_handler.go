package handler

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/taco-tortilla/notion-db-connector/internal/interfaces/request"
	"github.com/taco-tortilla/notion-db-connector/internal/usecase"
)

type FeedbackHandler interface {
	AddFeedback(*gin.Context)
}

type feedbackHandler struct {
	feedbackUseCase usecase.FeedbackUseCase
}

func NewFeedbackHandler(fu usecase.FeedbackUseCase) FeedbackHandler {
	return &feedbackHandler{
		feedbackUseCase: fu,
	}
}

func (fh feedbackHandler) AddFeedback(c *gin.Context) {
	request := request.FeedbackRequest{}
	if err := c.Bind(&request); err != nil {
		log.Fatal("Failed to bind request body")
	}

	c.Status(201)
}
