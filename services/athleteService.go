package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"ronin/interfaces"
	"ronin/models"
	"ronin/repositories"

	"strconv"

	"github.com/gorilla/mux"
)

var athleteRepo *repositories.AthleteRepository

func SetAthleteRepo(r *repositories.AthleteRepository) {
	athleteRepo = r
}

type AthleteUsername struct {
	Username string `json:"username" db:"username"`
}

type AthleteId struct {
	AthleteId int `json:"athleteId" db:"athlete_id"`
}

// athleteService implements the interfaces.AthleteService interface
type athleteService struct {
	repo *repositories.AthleteRepository
}

// NewAthleteService creates a new instance of AthleteService
func NewAthleteService(repo *repositories.AthleteRepository) interfaces.AthleteService {
	return &athleteService{
		repo: repo,
	}
}

func (s *athleteService) GetAll() ([]models.Athlete, error) {
	athletes, err := s.repo.GetAllAthletes()
	if err != nil {
		return nil, fmt.Errorf("failed to get all athletes: %w", err)
	}
	return athletes, nil
}

func (s *athleteService) GetByID(id string) (models.Athlete, error) {
	if id == "" {
		return models.Athlete{}, errors.New("athlete ID cannot be empty")
	}

	athlete, err := s.repo.GetAthleteById(id)
	if err != nil {
		return models.Athlete{}, fmt.Errorf("failed to get athlete by ID %s: %w", id, err)
	}
	return athlete, nil
}

func (s *athleteService) GetByUsername(username string) (models.Athlete, error) {
	if username == "" {
		return models.Athlete{}, errors.New("username cannot be empty")
	}

	athlete, err := s.repo.GetAthleteByUsername(username)
	if err != nil {
		return models.Athlete{}, fmt.Errorf("failed to get athlete by username %s: %w", username, err)
	}
	return athlete, nil
}

func (s *athleteService) Create(athlete models.Athlete) (int, error) {
	if err := s.validateAthlete(athlete); err != nil {
		return 0, fmt.Errorf("invalid athlete data: %w", err)
	}

	athleteID, err := s.repo.CreateAthlete(athlete)
	if err != nil {
		return 0, fmt.Errorf("failed to create athlete: %w", err)
	}
	return athleteID, nil
}

func (s *athleteService) Update(athlete models.Athlete) error {
	if err := s.validateAthlete(athlete); err != nil {
		return fmt.Errorf("invalid athlete data: %w", err)
	}

	if err := s.repo.UpdateAthlete(athlete); err != nil {
		return fmt.Errorf("failed to update athlete: %w", err)
	}
	return nil
}

func (s *athleteService) Delete(id string) error {
	if id == "" {
		return errors.New("athlete ID cannot be empty")
	}

	if err := s.repo.DeleteAthlete(id); err != nil {
		return fmt.Errorf("failed to delete athlete: %w", err)
	}
	return nil
}

func (s *athleteService) GetRecord(id string) (models.Record, error) {
	if id == "" {
		return models.Record{}, errors.New("athlete ID cannot be empty")
	}

	record, err := s.repo.GetAthleteRecord(id)
	if err != nil {
		return models.Record{}, fmt.Errorf("failed to get athlete record: %w", err)
	}

	// Convert AthleteRecord to Record
	return models.Record{
		Wins:   record.Wins,
		Losses: record.Losses,
		Draws:  record.Draws,
	}, nil
}

func (s *athleteService) GetAllUsernames() ([]string, error) {
	usernames, err := s.repo.GetAllUsernames()
	if err != nil {
		return nil, fmt.Errorf("failed to get all usernames: %w", err)
	}
	return usernames, nil
}

func (s *athleteService) AuthorizeUser(athlete models.Athlete) (bool, models.Athlete, error) {
	if athlete.Username == "" || athlete.Password == "" {
		return false, models.Athlete{}, errors.New("username and password are required")
	}

	isAuthorized, returnedAthlete, err := s.repo.IsAuthorizedUser(athlete)
	if err != nil {
		return false, models.Athlete{}, fmt.Errorf("failed to authorize user: %w", err)
	}
	return isAuthorized, returnedAthlete, nil
}

func (s *athleteService) FollowAthlete(follow models.Follow) error {
	if err := s.validateFollow(follow); err != nil {
		return fmt.Errorf("invalid follow data: %w", err)
	}

	if err := s.repo.FollowAthlete(follow); err != nil {
		return fmt.Errorf("failed to follow athlete: %w", err)
	}
	return nil
}

func (s *athleteService) UnfollowAthlete(followerID, followedID int) error {
	if followerID <= 0 || followedID <= 0 {
		return errors.New("invalid follower or followed ID")
	}

	if err := s.repo.UnfollowAthlete(followerID, followedID); err != nil {
		return fmt.Errorf("failed to unfollow athlete: %w", err)
	}
	return nil
}

func (s *athleteService) GetAthletesFollowed(id string) ([]models.Follow, error) {
	if id == "" {
		return nil, errors.New("athlete ID cannot be empty")
	}

	follows, err := s.repo.GetAthletesFollowed(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get followed athletes: %w", err)
	}

	// Convert []int to []models.Follow
	followList := make([]models.Follow, len(follows))
	followerID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid athlete ID format: %w", err)
	}

	for i, followedID := range follows {
		followList[i] = models.Follow{
			FollowerId: followerID,
			FollowedId: followedID,
		}
	}
	return followList, nil
}

func (s *athleteService) validateAthlete(athlete models.Athlete) error {
	if athlete.Username == "" {
		return errors.New("username is required")
	}
	if athlete.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func (s *athleteService) validateFollow(follow models.Follow) error {
	if follow.FollowerId <= 0 {
		return errors.New("invalid follower ID")
	}
	if follow.FollowedId <= 0 {
		return errors.New("invalid followed ID")
	}
	return nil
}

func GetAllAthleteUsernames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	usernames, err := athleteRepo.GetAllUsernames()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(&usernames)
}

func GetAllAthletes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	athletes, err := athleteRepo.GetAllAthletes()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(&athletes)
}

func GetAthlete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["athlete_id"]
	athletes, err := athleteRepo.GetAthleteById(id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(&athletes)
}

func GetAthleteByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	username := vars["username"]
	athlete, err := athleteRepo.GetAthleteByUsername(username)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(&athlete)
}

func IsAuthorizedUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Received authorization request")

	var athlete models.Athlete
	err := json.NewDecoder(r.Body).Decode(&athlete)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		log.Printf("Request body: %v", r.Body)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isAuthorized, returnedAthlete, err := athleteRepo.IsAuthorizedUser(athlete)
	if err != nil {
		log.Printf("Error in IsAuthorizedUser: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !isAuthorized {
		log.Println("User not authorized")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	} else if isAuthorized {
		idObj := AthleteId{AthleteId: returnedAthlete.AthleteId}
		log.Printf("Sending successful response: %+v", idObj)
		json.NewEncoder(w).Encode(&idObj)
	} else {
		log.Println("Sending false response")
		json.NewEncoder(w).Encode(false)
	}
}

func CreateAthlete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var athlete models.Athlete
	err := json.NewDecoder(r.Body).Decode(&athlete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	athleteId, err := athleteRepo.CreateAthlete(athlete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&athleteId)
}

func UpdateAthlete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var athlete models.Athlete
	err := json.NewDecoder(r.Body).Decode(&athlete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = athleteRepo.UpdateAthlete(athlete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Athlete updated successfully")
}

func DeleteAthlete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["athlete_id"]

	err := athleteRepo.DeleteAthlete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Athlete deleted successfully")
}

func GetAthleteRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["athlete_id"]

	record, err := athleteRepo.GetAthleteRecord(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&record)
}

func FollowAthlete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var follow models.Follow
	err := json.NewDecoder(r.Body).Decode(&follow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = athleteRepo.FollowAthlete(follow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Athlete followed successfully")
}

func UnfollowAthlete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	followerId, err := strconv.Atoi(vars["followerId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	followedId, err := strconv.Atoi(vars["followedId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = athleteRepo.UnfollowAthlete(followerId, followedId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Athlete unfollowed successfully")
}

func GetAthletesFollowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["athlete_id"]

	fmt.Println("Getting athletes followed by: ", id)

	follows, err := athleteRepo.GetAthletesFollowed(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&follows)
}
