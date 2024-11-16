package handler

import (
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
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if err := fh.feedbackUseCase.InsertFeedback(&request); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(201)
}
