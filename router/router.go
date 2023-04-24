package router

import (
	"github.com/gorilla/mux"

	"ronin/services"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/athletes", services.GetAllAthletes).Methods("GET")
	router.HandleFunc("/athlete/{athlete_id}", services.GetAthlete).Methods("GET")
	router.HandleFunc("/athlete", services.CreateAthlete).Methods("POST")
	router.HandleFunc("/athlete/{athlete_id}", services.UpdateAthlete).Methods("PUT")
	router.HandleFunc("/athlete/{athlete_id}", services.DeleteAthlete).Methods("DELETE")

	router.HandleFunc("/bouts", services.GetAllBouts).Methods("GET")
	router.HandleFunc("/bout/{bout_id}", services.GetBout).Methods("GET")
	router.HandleFunc("/bout", services.CreateBout).Methods("POST")
	router.HandleFunc("/bout/{bout_id}", services.UpdateBout).Methods("PUT")
	router.HandleFunc("/bout/{bout_id}", services.DeleteBout).Methods("DELETE")
	router.HandleFunc("/bout/{bout_id}/accept", services.AcceptBout).Methods("PUT")
	router.HandleFunc("/bout/{bout_id}/decline", services.DeclineBout).Methods("PUT")
	router.HandleFunc("/bout/{bout_id}/complete/{referee_id}", services.CompleteBout).Methods("PUT")
	router.HandleFunc("/bout/pending/{athlete_id}", services.GetPendingBouts).Methods("GET")
	router.HandleFunc("/bout/incomplete/{athlete_id}", services.GetIncompleteBouts).Methods("GET")

	router.HandleFunc("/gyms", services.GetAllGyms).Methods("GET")
	router.HandleFunc("/gym", services.CreateGym).Methods("POST")
	router.HandleFunc("/gym/{gym_id}", services.GetGym).Methods("GET")

	outcomeService := services.OutcomeService{}
	router.HandleFunc("/outcome", outcomeService.CreateOutcome).Methods("POST")
	router.HandleFunc("/outcome/{outcome_id}", services.GetOutcome).Methods("GET")
	router.HandleFunc("/outcome/bout/{bout_id}", services.GetOutcomeByBout).Methods("GET")
	router.HandleFunc("/outcome/bout/{bout_id}", outcomeService.CreateOutcomeByBout).Methods("POST")

	styleService := services.StyleService{}
	router.HandleFunc("/styles", services.GetAllStyles).Methods("GET")
	router.HandleFunc("/style", services.CreateStyle).Methods("POST")
	router.HandleFunc("/style/athlete/{athlete_id}", styleService.RegisterAthleteToStyle).Methods("POST")

	router.HandleFunc("/score/{athlete_id}", services.GetAthleteScore).Methods("GET")
	router.HandleFunc("/score/{athlete_id}/style/{style_id}", services.GetAthleteScoreByStyle).Methods("GET")

	return router
}
