package services

import (
	"encoding/json"
	"net/http"
	"ronin/interfaces"
	"ronin/models"

	"github.com/gorilla/mux"
)

type OutcomeHandler struct {
	service interfaces.OutcomeService
}

func NewOutcomeHandler(service interfaces.OutcomeService) *OutcomeHandler {
	return &OutcomeHandler{
		service: service,
	}
}

func (h *OutcomeHandler) GetAllOutcomes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	outcomes, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(outcomes)
}

func (h *OutcomeHandler) GetOutcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["outcome_id"]

	outcome, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(outcome)
}

func (h *OutcomeHandler) CreateOutcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var outcome models.Outcome
	if err := json.NewDecoder(r.Body).Decode(&outcome); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Create(outcome); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(outcome)
}

func (h *OutcomeHandler) GetOutcomeByBout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	boutID := vars["bout_id"]

	outcome, err := h.service.GetByBoutID(boutID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(outcome)
}

func (h *OutcomeHandler) CreateOutcomeByBout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	boutID := vars["bout_id"]

	var outcome models.Outcome
	if err := json.NewDecoder(r.Body).Decode(&outcome); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateForBout(outcome, boutID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(outcome)
}
