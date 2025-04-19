package interfaces

import "ronin/models"

// AthleteService defines the interface for athlete-related operations
type AthleteService interface {
	GetAll() ([]models.Athlete, error)
	GetByID(id string) (models.Athlete, error)
	GetByUsername(username string) (models.Athlete, error)
	Create(athlete models.Athlete) (int, error)
	Update(athlete models.Athlete) error
	Delete(id string) error
	GetRecord(id string) (models.Record, error)
	GetAllUsernames() ([]string, error)
	AuthorizeUser(athlete models.Athlete) (bool, models.Athlete, error)
	FollowAthlete(follow models.Follow) error
	UnfollowAthlete(followerID, followedID int) error
	GetAthletesFollowed(id string) ([]models.Follow, error)
}
