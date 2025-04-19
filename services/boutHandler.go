package services

import (
	"encoding/json"
	"log"
	"net/http"
	"ronin/interfaces"
	"ronin/models"

	"github.com/gorilla/mux"
)

// BoutHandler handles HTTP requests for bout operations
type BoutHandler struct {
	service interfaces.BoutService
}

// NewBoutHandler creates a new instance of BoutHandler
func NewBoutHandler(service interfaces.BoutService) *BoutHandler {
	return &BoutHandler{
		service: service,
	}
}

// GetAllBouts handles GET requests to retrieve all bouts
func (h *BoutHandler) GetAllBouts(w http.ResponseWriter, r *http.Request) {
	bouts, err := h.service.GetAll()
	if err != nil {
		SendError(w, "Failed to get bouts", http.StatusInternalServerError)
		return
	}
	SendJSON(w, bouts)
}

// GetBout handles GET requests to retrieve a specific bout
func (h *BoutHandler) GetBout(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bout_id"]
	if id == "" {
		SendError(w, "Invalid bout ID", http.StatusBadRequest)
		return
	}

	bout, err := h.service.GetByID(id)
	if err != nil {
		SendError(w, "Failed to get bout", http.StatusInternalServerError)
		return
	}
	SendJSON(w, bout)
}

// CreateBout handles POST requests to create a new bout
func (h *BoutHandler) CreateBout(w http.ResponseWriter, r *http.Request) {
	var bout models.Bout
	if err := json.NewDecoder(r.Body).Decode(&bout); err != nil {
		SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdBout, err := h.service.Create(bout)
	if err != nil {
		SendError(w, "Failed to create bout", http.StatusInternalServerError)
		return
	}
	SendJSON(w, createdBout)
}

// UpdateBout handles PUT requests to update an existing bout
func (h *BoutHandler) UpdateBout(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bout_id"]
	if id == "" {
		SendError(w, "Invalid bout ID", http.StatusBadRequest)
		return
	}

	var bout models.Bout
	if err := json.NewDecoder(r.Body).Decode(&bout); err != nil {
		SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.Update(id, bout); err != nil {
		SendError(w, "Failed to update bout", http.StatusInternalServerError)
		return
	}
	SendJSON(w, bout)
}

// DeleteBout handles DELETE requests to remove a bout
func (h *BoutHandler) DeleteBout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bout_id"]
	if id == "" {
		SendError(w, "Invalid bout ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(id); err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	SendJSON(w, id)
}

// AcceptBout handles PUT requests to accept a bout
func (h *BoutHandler) AcceptBout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bout_id"]
	if id == "" {
		SendError(w, "Invalid bout ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Accept(id); err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	SendJSON(w, id)
}

// DeclineBout handles PUT requests to decline a bout
func (h *BoutHandler) DeclineBout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bout_id"]
	if id == "" {
		SendError(w, "Invalid bout ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Decline(id); err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	SendJSON(w, id)
}

// CompleteBout handles PUT requests to complete a bout
func (h *BoutHandler) CompleteBout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boutId := vars["bout_id"]
	refereeId := vars["referee_id"]
	if boutId == "" || refereeId == "" {
		SendError(w, "Invalid bout or referee ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Complete(boutId, refereeId); err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	SendJSON(w, boutId)
}

// GetPendingBouts handles GET requests to retrieve pending bouts for an athlete
func (h *BoutHandler) GetPendingBouts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["athlete_id"]
	if id == "" {
		SendError(w, "Athlete ID is required", http.StatusBadRequest)
		return
	}

	bouts, err := h.service.GetPendingBouts(id)
	if err != nil {
		log.Printf("Error getting pending bouts for athlete %s: %v", id, err)
		SendError(w, "Failed to get pending bouts", http.StatusInternalServerError)
		return
	}

	if len(bouts) == 0 {
		SendJSON(w, []models.OutboundBout{})
		return
	}

	SendJSON(w, bouts)
}

// GetIncompleteBouts handles GET requests to retrieve incomplete bouts for an athlete
func (h *BoutHandler) GetIncompleteBouts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["athlete_id"]
	if id == "" {
		SendError(w, "Athlete ID is required", http.StatusBadRequest)
		return
	}

	bouts, err := h.service.GetIncompleteBouts(id)
	if err != nil {
		log.Printf("Error getting incomplete bouts for athlete %s: %v", id, err)
		SendError(w, "Failed to get incomplete bouts", http.StatusInternalServerError)
		return
	}

	if len(bouts) == 0 {
		SendJSON(w, []models.OutboundBout{})
		return
	}

	SendJSON(w, bouts)
}

// CancelBout handles PUT requests to cancel a bout
func (h *BoutHandler) CancelBout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boutId := vars["bout_id"]
	challengerId := vars["challenger_id"]
	if boutId == "" || challengerId == "" {
		SendError(w, "Invalid bout or challenger ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Cancel(boutId, challengerId); err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	SendJSON(w, boutId)
}
