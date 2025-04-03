package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"ronin/models"

	"strconv"

	"ronin/repositories"

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
	log.Printf("Decoded athlete data: %+v", athlete)

	isAuthorized, returnedAthlete, err := athleteRepo.IsAuthorizedUser(athlete)
	if err != nil {
		log.Printf("Error in IsAuthorizedUser: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Authorization result - isAuthorized: %v, athlete: %+v", isAuthorized, returnedAthlete)

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
	followerId, err := strconv.Atoi(vars["follower_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	followedId, err := strconv.Atoi(vars["followed_id"])
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

	follows, err := athleteRepo.GetAthletesFollowed(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&follows)
}
