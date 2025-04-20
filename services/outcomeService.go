package services

import (
	"errors"
	"fmt"
	"log"
	"ronin/interfaces"
	"ronin/models"
	"ronin/repositories"
)

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

func (s *outcomeService) CreateForBout(outcome models.Outcome, boutID string) error {
	log.Printf("Starting CreateForBout for bout %s with outcome: %+v", boutID, outcome)

	if err := s.validateOutcome(outcome); err != nil {
		log.Printf("Validation failed for outcome: %v", err)
		return fmt.Errorf("invalid outcome: %w", err)
	}

	if boutID == "" {
		log.Println("Bout ID is empty")
		return errors.New("bout ID cannot be empty")
	}

	// Check if bout exists and is in correct state
	bout, err := s.boutRepository.GetBoutById(boutID)
	if err != nil {
		log.Printf("Error getting bout %s: %v", boutID, err)
		return fmt.Errorf("failed to get bout: %w", err)
	}
	log.Printf("Found bout: %+v", bout)

	if !bout.Accepted {
		log.Printf("Bout %s is not accepted", boutID)
		return fmt.Errorf("bout %s is not accepted", boutID)
	}

	if bout.Completed {
		log.Printf("Bout %s is already completed", boutID)
		return fmt.Errorf("bout %s is already completed", boutID)
	}

	// Check if outcome already exists
	existingOutcome, err := s.outcomeRepo.GetOutcomeByBoutId(boutID)
	if err == nil {
		log.Printf("Outcome already exists for bout %s: %+v", boutID, existingOutcome)
		return fmt.Errorf("outcome already exists for bout %s", boutID)
	}

	// Create the outcome first
	createdOutcome, err := s.outcomeRepo.CreateOutcome(outcome)
	if err != nil {
		log.Printf("Failed to create outcome for bout %s: %v", boutID, err)
		return fmt.Errorf("failed to create outcome: %w", err)
	}
	log.Printf("Successfully created outcome with ID: %d", createdOutcome.OutcomeId)

	// Now update the athlete scores with the created outcome
	if err := s.updateAthleteScores(createdOutcome); err != nil {
		log.Printf("Failed to update athlete scores: %v", err)
		return fmt.Errorf("failed to update athlete scores: %w", err)
	}

	// Complete the bout
	if err := s.boutRepository.CompleteBoutByBoutId(boutID); err != nil {
		log.Printf("Failed to complete bout %s: %v", boutID, err)
		return fmt.Errorf("failed to complete bout: %w", err)
	}
	log.Printf("Successfully completed bout %s", boutID)

	log.Printf("Successfully created outcome for bout %s", boutID)
	return nil
}

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
