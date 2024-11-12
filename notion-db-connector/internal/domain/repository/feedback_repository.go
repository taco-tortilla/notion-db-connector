package repository

import "github.com/taco-tortilla/notion-db-connector/internal/domain"

type FeedbackRepository interface {
	Insert(fb *domain.Feedback) error
}
