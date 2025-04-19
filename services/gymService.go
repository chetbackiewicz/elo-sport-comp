package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ronin/models"
	"ronin/repositories"

	"github.com/gorilla/mux"
	// "github.com/jmoiron/sqlx"
)

var gymRepo *repositories.GymRepository

func SetGymRepo(r *repositories.GymRepository) {
	gymRepo = r
}

// GymService defines the interface for gym-related operations
type GymService interface {
	GetAll() ([]models.Gym, error)
	GetByID(id string) (models.Gym, error)
	Create(gym models.Gym) (models.Gym, error)
}

// gymService implements the GymService interface
type gymService struct {
	repo *repositories.GymRepository
}

// NewGymService creates a new instance of GymService
func NewGymService(repo *repositories.GymRepository) GymService {
	return &gymService{
		repo: repo,
	}
}

// GetAll retrieves all gyms
func (s *gymService) GetAll() ([]models.Gym, error) {
	gyms, err := s.repo.GetAllGyms()
	if err != nil {
		return nil, fmt.Errorf("failed to get all gyms: %w", err)
	}
	return gyms, nil
}

// GetByID retrieves a gym by its ID
func (s *gymService) GetByID(id string) (models.Gym, error) {
	if id == "" {
		return models.Gym{}, fmt.Errorf("gym ID cannot be empty")
	}

	gym, err := s.repo.GetGymById(id)
	if err != nil {
		return models.Gym{}, fmt.Errorf("failed to get gym by ID %s: %w", id, err)
	}
	return gym, nil
}

// Create creates a new gym
func (s *gymService) Create(gym models.Gym) (models.Gym, error) {
	if err := s.validateGym(gym); err != nil {
		return models.Gym{}, fmt.Errorf("invalid gym: %w", err)
	}

	createdGym, err := s.repo.CreateGym(gym)
	if err != nil {
		return models.Gym{}, fmt.Errorf("failed to create gym: %w", err)
	}
	return createdGym, nil
}

// validateGym validates the gym data
func (s *gymService) validateGym(gym models.Gym) error {
	if gym.Name == "" {
		return fmt.Errorf("gym name cannot be empty")
	}
	return nil
}

// GymHandler handles HTTP requests for gym operations
type GymHandler struct {
	service GymService
}

// NewGymHandler creates a new instance of GymHandler
func NewGymHandler(service GymService) *GymHandler {
	return &GymHandler{
		service: service,
	}
}

// GetAllGyms handles GET requests to retrieve all gyms
func (h *GymHandler) GetAllGyms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	gyms, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(gyms)
}

// GetGym handles GET requests to retrieve a specific gym
func (h *GymHandler) GetGym(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["gym_id"]

	gym, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(gym)
}

// CreateGym handles POST requests to create a new gym
func (h *GymHandler) CreateGym(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var gym models.Gym
	if err := json.NewDecoder(r.Body).Decode(&gym); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdGym, err := h.service.Create(gym)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(createdGym)
}
