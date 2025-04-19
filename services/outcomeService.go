package services

import (
	"errors"
	"fmt"
	"ronin/interfaces"
	"ronin/models"
	"ronin/repositories"
)

// OutcomeService defines the interface for outcome-related operations
type OutcomeService interface {
	GetAll() ([]models.Outcome, error)
	GetByID(id string) (models.Outcome, error)
	Create(outcome models.Outcome) error
	CreateForBout(outcome models.Outcome, boutID string) error
	GetByBoutID(boutID string) (models.Outcome, error)
}

// outcomeService implements the interfaces.OutcomeService interface
type outcomeService struct {
	outcomeRepo         *repositories.OutcomeRepository
	athleteScoreService *AthleteScoreService
	boutRepository      *repositories.BoutRepository
}

// NewOutcomeService creates a new instance of OutcomeService with all required dependencies
func NewOutcomeService(
	outcomeRepo *repositories.OutcomeRepository,
	athleteScoreService *AthleteScoreService,
	boutRepo *repositories.BoutRepository,
) interfaces.OutcomeService {
	return &outcomeService{
		outcomeRepo:         outcomeRepo,
		athleteScoreService: athleteScoreService,
		boutRepository:      boutRepo,
	}
}

// GetAll retrieves all outcomes
func (s *outcomeService) GetAll() ([]models.Outcome, error) {
	outcomes, err := s.outcomeRepo.GetAllOutcomes()
	if err != nil {
		return nil, fmt.Errorf("failed to get all outcomes: %w", err)
	}
	return outcomes, nil
}

// GetByID retrieves an outcome by its ID
func (s *outcomeService) GetByID(id string) (models.Outcome, error) {
	if id == "" {
		return models.Outcome{}, errors.New("outcome ID cannot be empty")
	}

	outcome, err := s.outcomeRepo.GetOutcomeById(id)
	if err != nil {
		return models.Outcome{}, fmt.Errorf("failed to get outcome by ID %s: %w", id, err)
	}
	return outcome, nil
}

// Create creates a new outcome and updates related athlete scores
func (s *outcomeService) Create(outcome models.Outcome) error {
	if err := s.validateOutcome(outcome); err != nil {
		return fmt.Errorf("invalid outcome: %w", err)
	}

	createdOutcome, err := s.outcomeRepo.CreateOutcome(outcome)
	if err != nil {
		return fmt.Errorf("failed to create outcome: %w", err)
	}

	if err := s.updateAthleteScores(createdOutcome); err != nil {
		return fmt.Errorf("failed to update athlete scores: %w", err)
	}

	return nil
}

// CreateForBout creates a new outcome for a specific bout and updates related athlete scores
func (s *outcomeService) CreateForBout(outcome models.Outcome, boutID string) error {
	if err := s.validateOutcome(outcome); err != nil {
		return fmt.Errorf("invalid outcome: %w", err)
	}

	if boutID == "" {
		return errors.New("bout ID cannot be empty")
	}

	if err := s.createOutcomeForBout(outcome, boutID); err != nil {
		return err
	}

	if err := s.updateAthleteScores(outcome); err != nil {
		return fmt.Errorf("failed to update athlete scores: %w", err)
	}

	return nil
}

// GetByBoutID retrieves an outcome associated with a specific bout
func (s *outcomeService) GetByBoutID(boutID string) (models.Outcome, error) {
	if boutID == "" {
		return models.Outcome{}, errors.New("bout ID cannot be empty")
	}

	outcome, err := s.outcomeRepo.GetOutcomeById(boutID) // Using GetOutcomeById as GetOutcomeByBoutId doesn't exist
	if err != nil {
		return models.Outcome{}, fmt.Errorf("failed to get outcome for bout %s: %w", boutID, err)
	}
	return outcome, nil
}

// validateOutcome validates the outcome data
func (s *outcomeService) validateOutcome(outcome models.Outcome) error {
	if outcome.WinnerId == 0 {
		return errors.New("winner ID is required")
	}
	if outcome.LoserId == 0 {
		return errors.New("loser ID is required")
	}
	if outcome.StyleId == 0 {
		return errors.New("style ID is required")
	}
	return nil
}

// createOutcomeForBout handles the creation of an outcome for a specific bout
func (s *outcomeService) createOutcomeForBout(outcome models.Outcome, boutID string) error {
	var err error
	if outcome.IsDraw {
		err = s.outcomeRepo.CreateOutcomeByBoutIdDraw(&outcome, boutID)
	} else {
		err = s.outcomeRepo.CreateOutcomeByBoutIdNotDraw(&outcome, boutID)
	}

	if err != nil {
		return fmt.Errorf("failed to create outcome for bout %s: %w", boutID, err)
	}

	return nil
}

// updateAthleteScores updates the scores for both athletes involved in the outcome
func (s *outcomeService) updateAthleteScores(outcome models.Outcome) error {
	loserScore, err := s.athleteScoreService.GetAthleteScoreByStyle(outcome.LoserId, outcome.StyleId)
	if err != nil {
		return fmt.Errorf("failed to get loser score: %w", err)
	}

	winnerScore, err := s.athleteScoreService.GetAthleteScoreByStyle(outcome.WinnerId, outcome.StyleId)
	if err != nil {
		return fmt.Errorf("failed to get winner score: %w", err)
	}

	err = s.athleteScoreService.CalculateNewScores(&winnerScore, &loserScore, outcome.IsDraw, outcome.OutcomeId)
	if err != nil {
		return fmt.Errorf("failed to calculate new scores: %w", err)
	}

	return nil
}
