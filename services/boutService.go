package services

import (
	"errors"
	"fmt"
	"log"
	"ronin/interfaces"
	"ronin/models"
	"ronin/repositories"
)

// boutService implements the interfaces.BoutService interface
type boutService struct {
	repo *repositories.BoutRepository
}

// NewBoutService creates a new instance of BoutService
func NewBoutService(repo *repositories.BoutRepository) interfaces.BoutService {
	return &boutService{
		repo: repo,
	}
}

// GetAll retrieves all bouts
func (s *boutService) GetAll() ([]models.OutboundBout, error) {
	log.Println("Getting all bouts")
	// First get all bout IDs
	bouts, err := s.repo.GetAllBouts()
	if err != nil {
		log.Printf("Error getting all bouts: %v", err)
		return nil, fmt.Errorf("failed to get all bouts: %w", err)
	}
	log.Printf("Found %d bouts", len(bouts))

	// Convert each bout to outbound bout
	outboundBouts := make([]models.OutboundBout, 0, len(bouts))
	for _, bout := range bouts {
		log.Printf("Converting bout ID %d to outbound bout", bout.BoutId)
		outboundBout, err := s.repo.GetOutboundBoutByBoutId(bout.BoutId)
		if err != nil {
			log.Printf("Error getting outbound bout for bout ID %d: %v", bout.BoutId, err)
			return nil, fmt.Errorf("failed to get outbound bout for bout ID %d: %w", bout.BoutId, err)
		}
		outboundBouts = append(outboundBouts, outboundBout)
	}

	log.Printf("Successfully retrieved %d outbound bouts", len(outboundBouts))
	return outboundBouts, nil
}

// GetByID retrieves a bout by its ID
func (s *boutService) GetByID(id string) (models.OutboundBout, error) {
	log.Printf("Getting bout by ID: %s", id)
	if id == "" {
		log.Println("Bout ID cannot be empty")
		return models.OutboundBout{}, errors.New("bout ID cannot be empty")
	}

	// First get the bout to verify it exists and get its ID
	bout, err := s.repo.GetBoutById(id)
	if err != nil {
		log.Printf("Error getting bout by ID %s: %v", id, err)
		return models.OutboundBout{}, fmt.Errorf("failed to get bout by ID %s: %w", id, err)
	}
	log.Printf("Found bout with ID %s", id)

	// Then get the outbound bout
	outboundBout, err := s.repo.GetOutboundBoutByBoutId(bout.BoutId)
	if err != nil {
		log.Printf("Error getting outbound bout for bout ID %s: %v", id, err)
		return models.OutboundBout{}, fmt.Errorf("failed to get outbound bout for bout ID %s: %w", id, err)
	}

	log.Printf("Successfully retrieved outbound bout for ID %s", id)
	return outboundBout, nil
}

// Create creates a new bout
func (s *boutService) Create(bout models.Bout) (models.OutboundBout, error) {
	if err := s.validateBout(bout); err != nil {
		return models.OutboundBout{}, fmt.Errorf("invalid bout: %w", err)
	}

	boutID, err := s.repo.CreateBout(bout)
	if err != nil {
		return models.OutboundBout{}, fmt.Errorf("failed to create bout: %w", err)
	}

	createdBout, err := s.repo.GetOutboundBoutByBoutId(boutID)
	if err != nil {
		return models.OutboundBout{}, fmt.Errorf("failed to get created bout: %w", err)
	}

	return createdBout, nil
}

// Update updates an existing bout
func (s *boutService) Update(id string, bout models.Bout) error {
	if id == "" {
		return errors.New("bout ID cannot be empty")
	}

	if err := s.validateBout(bout); err != nil {
		return fmt.Errorf("invalid bout: %w", err)
	}

	if err := s.repo.UpdateBout(id, bout); err != nil {
		return fmt.Errorf("failed to update bout: %w", err)
	}

	return nil
}

// Delete deletes a bout
func (s *boutService) Delete(id string) error {
	if id == "" {
		return errors.New("bout ID cannot be empty")
	}

	if err := s.repo.DeleteBout(id); err != nil {
		return fmt.Errorf("failed to delete bout: %w", err)
	}

	return nil
}

// Accept accepts a bout
func (s *boutService) Accept(id string) error {
	if id == "" {
		return errors.New("bout ID cannot be empty")
	}

	if err := s.repo.AcceptBout(id); err != nil {
		return fmt.Errorf("failed to accept bout: %w", err)
	}

	return nil
}

// Decline declines a bout
func (s *boutService) Decline(id string) error {
	if id == "" {
		return errors.New("bout ID cannot be empty")
	}

	if err := s.repo.DeclineBout(id); err != nil {
		return fmt.Errorf("failed to decline bout: %w", err)
	}

	return nil
}

// Complete completes a bout
func (s *boutService) Complete(boutID string, refereeID string) error {
	if boutID == "" {
		return errors.New("bout ID cannot be empty")
	}
	if refereeID == "" {
		return errors.New("referee ID cannot be empty")
	}

	if err := s.repo.CompleteBout(boutID, refereeID); err != nil {
		return fmt.Errorf("failed to complete bout: %w", err)
	}

	return nil
}

// Cancel cancels a bout
func (s *boutService) Cancel(boutID string, challengerID string) error {
	if boutID == "" {
		return errors.New("bout ID cannot be empty")
	}
	if challengerID == "" {
		return errors.New("challenger ID cannot be empty")
	}

	if err := s.repo.CancelBout(boutID, challengerID); err != nil {
		return fmt.Errorf("failed to cancel bout: %w", err)
	}

	return nil
}

// GetPendingBouts retrieves all pending bouts for an athlete
func (s *boutService) GetPendingBouts(athleteID string) ([]models.OutboundBout, error) {
	log.Printf("Getting pending bouts for athlete ID: %s", athleteID)
	if athleteID == "" {
		log.Println("Athlete ID cannot be empty")
		return nil, errors.New("athlete ID cannot be empty")
	}

	bouts, err := s.repo.GetPendingBoutsByAthleteId(athleteID)
	if err != nil {
		log.Printf("Error getting pending bouts for athlete %s: %v", athleteID, err)
		return nil, fmt.Errorf("failed to get pending bouts: %w", err)
	}

	log.Printf("Found %d pending bouts for athlete %s", len(bouts), athleteID)
	for i, bout := range bouts {
		log.Printf("Pending bout %d: ID=%d, Challenger=%s %s (ID=%d), Acceptor=%s %s (ID=%d), Style=%s (ID=%d)",
			i+1, bout.BoutId,
			bout.ChallengerFirstName, bout.ChallengerLastName, bout.ChallengerId,
			bout.AcceptorFirstName, bout.AcceptorLastName, bout.AcceptorId,
			bout.Style, bout.StyleId)
	}

	return bouts, nil
}

// GetIncompleteBouts retrieves all incomplete bouts for an athlete
func (s *boutService) GetIncompleteBouts(athleteID string) ([]models.OutboundBout, error) {
	log.Printf("Getting incomplete bouts for athlete ID: %s", athleteID)
	if athleteID == "" {
		log.Println("Athlete ID cannot be empty")
		return nil, errors.New("athlete ID cannot be empty")
	}

	bouts, err := s.repo.GetIncompleteBoutsByAthleteId(athleteID)
	if err != nil {
		log.Printf("Error getting incomplete bouts for athlete %s: %v", athleteID, err)
		return nil, fmt.Errorf("failed to get incomplete bouts: %w", err)
	}

	log.Printf("Found %d incomplete bouts for athlete %s", len(bouts), athleteID)
	for i, bout := range bouts {
		log.Printf("Incomplete bout %d: ID=%d, Challenger=%s %s (ID=%d), Acceptor=%s %s (ID=%d), Style=%s (ID=%d)",
			i+1, bout.BoutId,
			bout.ChallengerFirstName, bout.ChallengerLastName, bout.ChallengerId,
			bout.AcceptorFirstName, bout.AcceptorLastName, bout.AcceptorId,
			bout.Style, bout.StyleId)
	}

	return bouts, nil
}

// validateBout validates the bout data
func (s *boutService) validateBout(bout models.Bout) error {
	if bout.ChallengerId == 0 {
		return errors.New("challenger ID is required")
	}
	if bout.AcceptorId == 0 {
		return errors.New("acceptor ID is required")
	}
	if bout.ChallengerId == bout.AcceptorId {
		return errors.New("challenger and acceptor cannot be the same athlete")
	}
	return nil
}
