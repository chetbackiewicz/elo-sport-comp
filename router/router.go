package router

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"ronin/services"
)

const base_url = "/api/v1"

var (
	outcomeHandler      *services.OutcomeHandler
	athleteHandler      *services.AthleteHandler
	athleteScoreHandler *services.AthleteScoreHandler
	boutHandler         *services.BoutHandler
	feedHandler         *services.FeedHandler
	gymHandler          *services.GymHandler
	styleHandler        *services.StyleHandler
)

func SetOutcomeHandler(h *services.OutcomeHandler) {
	outcomeHandler = h
}

func SetAthleteHandler(h *services.AthleteHandler) {
	athleteHandler = h
}

func SetAthleteScoreHandler(h *services.AthleteScoreHandler) {
	athleteScoreHandler = h
}

func SetBoutHandler(h *services.BoutHandler) {
	boutHandler = h
}

func SetFeedHandler(h *services.FeedHandler) {
	feedHandler = h
}

func SetGymHandler(h *services.GymHandler) {
	gymHandler = h
}

func SetStyleHandler(h *services.StyleHandler) {
	styleHandler = h
}

// LoggingMiddleware logs all incoming requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// Create a custom response writer to capture the status code
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Call the next handler
		next.ServeHTTP(rw, r)

		// Log the response
		duration := time.Since(start)
		log.Printf("Completed %s %s in %v with status %d",
			r.Method, r.URL.Path, duration, rw.statusCode)
	})
}

// Custom response writer to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	// Apply logging middleware to all routes
	router.Use(LoggingMiddleware)

	// Athlete routes
	router.HandleFunc(base_url+"/athletes", athleteHandler.GetAllAthletes).Methods("GET")
	router.HandleFunc(base_url+"/athlete/{athlete_id}", athleteHandler.GetAthlete).Methods("GET")
	router.HandleFunc(base_url+"/athlete", athleteHandler.CreateAthlete).Methods("POST")
	router.HandleFunc(base_url+"/athlete/{athlete_id}", athleteHandler.UpdateAthlete).Methods("PUT")
	router.HandleFunc(base_url+"/athlete/{athlete_id}", athleteHandler.DeleteAthlete).Methods("DELETE")
	router.HandleFunc(base_url+"/athlete/all/usernames", athleteHandler.GetAllAthleteUsernames).Methods("GET")
	router.HandleFunc(base_url+"/athlete/{athlete_id}/record", athleteHandler.GetAthleteRecord).Methods("GET")
	router.HandleFunc(base_url+"/athlete/authorize", athleteHandler.IsAuthorizedUser).Methods("POST")
	router.HandleFunc(base_url+"/athletes/follow", athleteHandler.FollowAthlete).Methods("POST")
	router.HandleFunc(base_url+"/athletes/{followerId}/{followedId}/unfollow", athleteHandler.UnfollowAthlete).Methods("DELETE")
	router.HandleFunc(base_url+"/athletes/following/{id}", athleteHandler.GetAthletesFollowed).Methods("GET")

	// Bout routes
	router.HandleFunc(base_url+"/bouts", boutHandler.GetAllBouts).Methods("GET")
	router.HandleFunc(base_url+"/bout/{bout_id}", boutHandler.GetBout).Methods("GET")
	router.HandleFunc(base_url+"/bout", boutHandler.CreateBout).Methods("POST")
	router.HandleFunc(base_url+"/bout/{bout_id}", boutHandler.UpdateBout).Methods("PUT")
	router.HandleFunc(base_url+"/bout/{bout_id}", boutHandler.DeleteBout).Methods("DELETE")
	router.HandleFunc(base_url+"/bout/{bout_id}/accept", boutHandler.AcceptBout).Methods("PUT")
	router.HandleFunc(base_url+"/bout/{bout_id}/decline", boutHandler.DeclineBout).Methods("PUT")
	router.HandleFunc(base_url+"/bout/{bout_id}/complete/{referee_id}", boutHandler.CompleteBout).Methods("PUT")
	router.HandleFunc(base_url+"/bout/cancel/{bout_id}/{challenger_id}", boutHandler.CancelBout).Methods("PUT")
	router.HandleFunc(base_url+"/bouts/pending/{athlete_id}", boutHandler.GetPendingBouts).Methods("GET")
	router.HandleFunc(base_url+"/bouts/incomplete/{athlete_id}", boutHandler.GetIncompleteBouts).Methods("GET")

	// Outcome routes
	router.HandleFunc(base_url+"/outcomes", outcomeHandler.GetAllOutcomes).Methods("GET")
	router.HandleFunc(base_url+"/outcome/{outcome_id}", outcomeHandler.GetOutcome).Methods("GET")
	router.HandleFunc(base_url+"/outcome", outcomeHandler.CreateOutcome).Methods("POST")
	router.HandleFunc(base_url+"/outcome/bout/{bout_id}", outcomeHandler.GetOutcomeByBout).Methods("GET")
	router.HandleFunc(base_url+"/outcome/bout/{bout_id}", outcomeHandler.CreateOutcomeByBout).Methods("POST")

	// Style routes
	router.HandleFunc(base_url+"/styles", styleHandler.GetAllStyles).Methods("GET")
	router.HandleFunc(base_url+"/style", styleHandler.CreateStyle).Methods("POST")
	router.HandleFunc(base_url+"/style/athlete/{athlete_id}", styleHandler.RegisterAthleteToStyle).Methods("POST")
	router.HandleFunc(base_url+"/styles/athlete/{athlete_id}", styleHandler.RegisterMultipleStylesToAthlete).Methods("POST")
	router.HandleFunc(base_url+"/styles/common/{athlete_id}/{challenger_id}", styleHandler.GetCommonStyles).Methods("GET")

	// Athlete Score routes
	router.HandleFunc(base_url+"/score/{athlete_id}", athleteScoreHandler.GetAthleteScore).Methods("GET")
	router.HandleFunc(base_url+"/score/{athlete_id}/all", athleteScoreHandler.GetAthleteScore).Methods("GET")
	router.HandleFunc(base_url+"/score/{athlete_id}/style/{style_id}", athleteScoreHandler.GetAthleteScoreByStyle).Methods("GET")
	router.HandleFunc(base_url+"/score/{athlete_id}/style/{style_id}/history", athleteScoreHandler.GetAthleteScoreHistory).Methods("GET")

	// Feed routes
	router.HandleFunc(base_url+"/feed/{athlete_id}", feedHandler.GetFeedByAthleteID).Methods("GET")

	// Gym routes
	router.HandleFunc(base_url+"/gyms", gymHandler.GetAllGyms).Methods("GET")
	router.HandleFunc(base_url+"/gym/{gym_id}", gymHandler.GetGym).Methods("GET")
	router.HandleFunc(base_url+"/gym", gymHandler.CreateGym).Methods("POST")

	return router
}
