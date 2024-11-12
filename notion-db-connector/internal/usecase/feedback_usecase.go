package usecase

import (
	"github.com/taco-tortilla/notion-db-connector/internal/domain"
	"github.com/taco-tortilla/notion-db-connector/internal/domain/repository"
	"github.com/taco-tortilla/notion-db-connector/internal/interfaces/request"
)

type FeedbackUseCase interface {
	InsertFeedback(fr *request.FeedbackRequest) error
	Insert(fb *domain.Feedback) error
}

type feedbackUseCase struct {
	feedbackRepository repository.FeedbackRepository
}

func NewFeedbackUseCase(fr repository.FeedbackRepository) FeedbackUseCase {
	return &feedbackUseCase{
		feedbackRepository: fr,
	}
}

func (fu feedbackUseCase) InsertFeedback(fr *request.FeedbackRequest) error {
	feedback, err := domain.NewFeedback(fr.CompanyName, fr.ContractType, fr.Plan, fr.CreatedBy, fr.Body)
	if err != nil {
		return err
	}
	return fu.Insert(feedback)
}

func (fu feedbackUseCase) Insert(fb *domain.Feedback) error {
	return fu.feedbackRepository.Insert(fb)
}
