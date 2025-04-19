package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"ronin/models"
	"ronin/repositories"

	"github.com/gorilla/mux"
)

// FeedService defines the interface for feed-related operations
type FeedService interface {
	GetByAthleteID(athleteID string) ([]models.Feed, error)
}

// feedService implements the FeedService interface
type feedService struct {
	repo *repositories.FeedRepository
}

// NewFeedService creates a new instance of FeedService
func NewFeedService(repo *repositories.FeedRepository) FeedService {
	return &feedService{
		repo: repo,
	}
}

// GetByAthleteID retrieves feed items for a specific athlete
func (s *feedService) GetByAthleteID(athleteID string) ([]models.Feed, error) {
	if athleteID == "" {
		return nil, fmt.Errorf("athlete ID cannot be empty")
	}

	feed, err := s.repo.GetFeedByAthleteId(athleteID)
	if err == sql.ErrNoRows {
		return []models.Feed{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get feed for athlete %s: %w", athleteID, err)
	}

	return feed, nil
}

// FeedHandler handles HTTP requests for feed operations
type FeedHandler struct {
	service FeedService
}

// NewFeedHandler creates a new instance of FeedHandler
func NewFeedHandler(service FeedService) *FeedHandler {
	return &FeedHandler{
		service: service,
	}
}

// GetFeedByAthleteID handles GET requests to retrieve feed items for an athlete
func (h *FeedHandler) GetFeedByAthleteID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	athleteID := vars["athlete_id"]

	feed, err := h.service.GetByAthleteID(athleteID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(feed) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	json.NewEncoder(w).Encode(feed)
}
