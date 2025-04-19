package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ronin/interfaces"
	"ronin/models"

	"github.com/gorilla/mux"
)

type AthleteHandler struct {
	service interfaces.AthleteService
}

func NewAthleteHandler(service interfaces.AthleteService) *AthleteHandler {
	return &AthleteHandler{
		service: service,
	}
}

func (h *AthleteHandler) GetAllAthletes(w http.ResponseWriter, r *http.Request) {
	athletes, err := h.service.GetAll()
	if err != nil {
		SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SendJSON(w, athletes)
}

func (h *AthleteHandler) GetAthlete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	athlete, err := h.service.GetByID(id)
	if err != nil {
		SendError(w, err.Error(), http.StatusNotFound)
		return
	}
	SendJSON(w, athlete)
}

func (h *AthleteHandler) GetAthleteByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	athlete, err := h.service.GetByUsername(username)
	if err != nil {
		SendError(w, err.Error(), http.StatusNotFound)
		return
	}
	SendJSON(w, athlete)
}

func (h *AthleteHandler) CreateAthlete(w http.ResponseWriter, r *http.Request) {
	var athlete models.Athlete
	if err := json.NewDecoder(r.Body).Decode(&athlete); err != nil {
		SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.service.Create(athlete)
	if err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	athlete.AthleteId = id
	SendJSON(w, athlete)
}

func (h *AthleteHandler) UpdateAthlete(w http.ResponseWriter, r *http.Request) {
	var athlete models.Athlete
	if err := json.NewDecoder(r.Body).Decode(&athlete); err != nil {
		SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.Update(athlete); err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	SendJSON(w, athlete)
}

func (h *AthleteHandler) DeleteAthlete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.service.Delete(id); err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	SendJSON(w, map[string]string{"message": "Athlete deleted successfully"})
}

func (h *AthleteHandler) GetAthleteRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	record, err := h.service.GetRecord(id)
	if err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	SendJSON(w, record)
}

func (h *AthleteHandler) GetAllAthleteUsernames(w http.ResponseWriter, r *http.Request) {
	usernames, err := h.service.GetAllUsernames()
	if err != nil {
		SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SendJSON(w, usernames)
}

func (h *AthleteHandler) IsAuthorizedUser(w http.ResponseWriter, r *http.Request) {
	var athlete models.Athlete
	if err := json.NewDecoder(r.Body).Decode(&athlete); err != nil {
		SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if athlete.Username == "" || athlete.Password == "" {
		SendError(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	isAuthorized, returnedAthlete, err := h.service.AuthorizeUser(athlete)
	if err != nil {
		SendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if !isAuthorized {
		SendError(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	SendJSON(w, map[string]interface{}{
		"athleteId": returnedAthlete.AthleteId,
		"success":   true,
	})
}

func (h *AthleteHandler) FollowAthlete(w http.ResponseWriter, r *http.Request) {
	var follow models.Follow
	if err := json.NewDecoder(r.Body).Decode(&follow); err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.FollowAthlete(follow); err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	SendJSON(w, map[string]string{"message": "Athlete followed successfully"})
}

func (h *AthleteHandler) UnfollowAthlete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	followerID, err := strconv.Atoi(vars["followerId"])
	if err != nil {
		SendError(w, "Invalid follower ID", http.StatusBadRequest)
		return
	}

	followedID, err := strconv.Atoi(vars["followedId"])
	if err != nil {
		SendError(w, "Invalid followed ID", http.StatusBadRequest)
		return
	}

	if err := h.service.UnfollowAthlete(followerID, followedID); err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	SendJSON(w, map[string]string{"message": "Athlete unfollowed successfully"})
}
