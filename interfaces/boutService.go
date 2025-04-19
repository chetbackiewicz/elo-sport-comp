package interfaces

import "ronin/models"

// BoutService defines the interface for bout-related operations
type BoutService interface {
	GetAll() ([]models.OutboundBout, error)
	GetByID(id string) (models.OutboundBout, error)
	Create(bout models.Bout) (models.OutboundBout, error)
	Update(id string, bout models.Bout) error
	Delete(id string) error
	Accept(id string) error
	Decline(id string) error
	Complete(boutID string, refereeID string) error
	Cancel(boutID string, challengerID string) error
	GetPendingBouts(athleteID string) ([]models.OutboundBout, error)
	GetIncompleteBouts(athleteID string) ([]models.OutboundBout, error)
}
