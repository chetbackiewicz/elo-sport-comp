package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ronin/models"
	"ronin/repositories"
	"strconv"

	"github.com/gorilla/mux"
)

var styleRepo *repositories.StyleRepository

func SetStyleRepo(r *repositories.StyleRepository) {
	styleRepo = r
}

// StyleService defines the interface for style-related operations
type StyleService interface {
	GetAll() ([]models.Style, error)
	Create(style models.Style) error
	RegisterAthleteToStyle(athleteID int, styleID int) error
	RegisterMultipleStylesToAthlete(athleteID int, styles []int) error
	GetCommonStyles(acceptorID, challengerID string) ([]models.Style, error)
}

// styleService implements the StyleService interface
type styleService struct {
	repo                *repositories.StyleRepository
	athleteScoreService *AthleteScoreService
}

// NewStyleService creates a new instance of StyleService
func NewStyleService(repo *repositories.StyleRepository, athleteScoreService *AthleteScoreService) StyleService {
	return &styleService{
		repo:                repo,
		athleteScoreService: athleteScoreService,
	}
}

// GetAll retrieves all styles
func (s *styleService) GetAll() ([]models.Style, error) {
	styles, err := s.repo.GetAllStyles()
	if err != nil {
		return nil, fmt.Errorf("failed to get all styles: %w", err)
	}
	return styles, nil
}

// Create creates a new style
func (s *styleService) Create(style models.Style) error {
	if err := s.validateStyle(style); err != nil {
		return fmt.Errorf("invalid style: %w", err)
	}

	if err := s.repo.CreateStyle(style); err != nil {
		return fmt.Errorf("failed to create style: %w", err)
	}
	return nil
}

// RegisterAthleteToStyle registers an athlete to a specific style
func (s *styleService) RegisterAthleteToStyle(athleteID int, styleID int) error {
	if err := s.repo.RegisterAthleteToStyle(athleteID, styleID); err != nil {
		return fmt.Errorf("failed to register athlete %d to style %d: %w", athleteID, styleID, err)
	}

	if err := s.athleteScoreService.CreateAthleteScoreUponRegistration(athleteID, styleID); err != nil {
		return fmt.Errorf("failed to create athlete score: %w", err)
	}

	return nil
}

// RegisterMultipleStylesToAthlete registers multiple styles to an athlete
func (s *styleService) RegisterMultipleStylesToAthlete(athleteID int, styles []int) error {
	for _, styleID := range styles {
		if err := s.RegisterAthleteToStyle(athleteID, styleID); err != nil {
			return fmt.Errorf("failed to register style %d: %w", styleID, err)
		}
	}
	return nil
}

// GetCommonStyles retrieves styles common between two athletes
func (s *styleService) GetCommonStyles(acceptorID, challengerID string) ([]models.Style, error) {
	if acceptorID == "" || challengerID == "" {
		return nil, fmt.Errorf("both acceptor and challenger IDs are required")
	}

	styles, err := s.repo.GetCommonStyles(acceptorID, challengerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get common styles: %w", err)
	}
	return styles, nil
}

// validateStyle validates the style data
func (s *styleService) validateStyle(style models.Style) error {
	if style.StyleName == "" {
		return fmt.Errorf("style name cannot be empty")
	}
	return nil
}

// StyleHandler handles HTTP requests for style operations
type StyleHandler struct {
	service StyleService
}

// NewStyleHandler creates a new instance of StyleHandler
func NewStyleHandler(service StyleService) *StyleHandler {
	return &StyleHandler{
		service: service,
	}
}

// GetAllStyles handles GET requests to retrieve all styles
func (h *StyleHandler) GetAllStyles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	styles, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(styles)
}

// CreateStyle handles POST requests to create a new style
func (h *StyleHandler) CreateStyle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var style models.Style
	if err := json.NewDecoder(r.Body).Decode(&style); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Create(style); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(style)
}

// RegisterAthleteToStyle handles POST requests to register an athlete to a style
func (h *StyleHandler) RegisterAthleteToStyle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	athleteIDStr := vars["athlete_id"]

	athleteID, err := strconv.Atoi(athleteIDStr)
	if err != nil {
		http.Error(w, "invalid athlete ID", http.StatusBadRequest)
		return
	}

	var style models.Style
	if err := json.NewDecoder(r.Body).Decode(&style); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.RegisterAthleteToStyle(athleteID, style.StyleId); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(style)
}

// RegisterMultipleStylesToAthlete handles POST requests to register multiple styles to an athlete
func (h *StyleHandler) RegisterMultipleStylesToAthlete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request models.RegisterStylesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.RegisterMultipleStylesToAthlete(request.AthleteID, request.Styles); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetCommonStyles handles GET requests to retrieve common styles between two athletes
func (h *StyleHandler) GetCommonStyles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	acceptorID := vars["athlete_id"]
	challengerID := vars["challenger_id"]

	styles, err := h.service.GetCommonStyles(acceptorID, challengerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(styles)
}
