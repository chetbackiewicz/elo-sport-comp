package interfaces

import "ronin/models"


type StyleService interface {
	GetAll() ([]models.Style, error)
	Create(style models.Style) error
	RegisterAthleteToStyle(athleteID int, styleID int) error
	RegisterMultipleStylesToAthlete(athleteID int, styles []int) error
	GetCommonStyles(acceptorID, challengerID string) ([]models.Style, error)
}
