package interfaces

import "ronin/models"

// GymService defines the interface for gym-related operations
type GymService interface {
	GetAll() ([]models.Gym, error)
	GetByID(id string) (models.Gym, error)
	Create(gym models.Gym) (models.Gym, error)
}
