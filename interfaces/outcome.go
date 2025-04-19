package interfaces

import "ronin/models"

// OutcomeService defines the interface for outcome-related operations
type OutcomeService interface {
	GetAll() ([]models.Outcome, error)
	GetByID(id string) (models.Outcome, error)
	Create(outcome models.Outcome) error
	CreateForBout(outcome models.Outcome, boutID string) error
	GetByBoutID(boutID string) (models.Outcome, error)
}
