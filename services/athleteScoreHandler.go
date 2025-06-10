package services

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AthleteScoreHandler struct {
	service *AthleteScoreService
}

func NewAthleteScoreHandler(service *AthleteScoreService) *AthleteScoreHandler {
	return &AthleteScoreHandler{
		service: service,
	}
}

func (h *AthleteScoreHandler) GetAthleteScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idStr := vars["athlete_id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid athlete_id: %v\n", err)
		http.Error(w, "Invalid athlete_id", http.StatusBadRequest)
		return
	}

	athleteScores, err := h.service.GetAllAthleteScoresByAthleteId(id)
	if err != nil {
		log.Printf("Error getting athlete scores: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(athleteScores); err != nil {
		log.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (h *AthleteScoreHandler) GetAthleteScoreByStyle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idStr := vars["athlete_id"]
	styleStr := vars["style_id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid athlete_id: %v\n", err)
		http.Error(w, "Invalid athlete_id", http.StatusBadRequest)
		return
	}

	style, err := strconv.Atoi(styleStr)
	if err != nil {
		log.Printf("Invalid style_id: %v\n", err)
		http.Error(w, "Invalid style_id", http.StatusBadRequest)
		return
	}

	athleteScore, err := h.service.GetAthleteScoreByStyle(id, style)
	if err != nil {
		log.Printf("Error getting athlete score: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(athleteScore); err != nil {
		log.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (h *AthleteScoreHandler) GetAthleteScoreHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	athleteIdStr := vars["athlete_id"]
	styleIdStr := vars["style_id"]

	athleteId, err := strconv.Atoi(athleteIdStr)
	if err != nil {
		log.Printf("Invalid athlete_id: %v\n", err)
		http.Error(w, "Invalid athlete_id", http.StatusBadRequest)
		return
	}

	styleId, err := strconv.Atoi(styleIdStr)
	if err != nil {
		log.Printf("Invalid style_id: %v\n", err)
		http.Error(w, "Invalid style_id", http.StatusBadRequest)
		return
	}

	history, err := h.service.GetAthleteScoreHistory(athleteId, styleId)
	if err != nil {
		log.Printf("Error getting athlete score history: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(history); err != nil {
		log.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
