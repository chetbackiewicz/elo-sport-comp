package interfaces

import "ronin/models"

// FeedService defines the interface for feed-related operations
type FeedService interface {
	GetByAthleteID(athleteID string) ([]models.Feed, error)
}
