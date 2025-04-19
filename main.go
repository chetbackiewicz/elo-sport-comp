package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ronin/interfaces"
	"ronin/repositories"
	"ronin/router"
	"ronin/services"
	"ronin/utils"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("In Main App")

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get database connection
	dbconn := utils.GetConnection()

	// Initialize repositories
	athleteRepo := repositories.NewAthleteRepository(dbconn)
	outcomeRepo := repositories.NewOutcomeRepository(dbconn)
	boutRepo := repositories.NewBoutRepository(dbconn)
	athleteScoreRepo := repositories.NewAthleteScoreRepository(dbconn)
	feedRepo := repositories.NewFeedRepository(dbconn)
	gymRepo := repositories.NewGymRepository(dbconn)
	styleRepo := repositories.NewStyleRepository(dbconn)

	// Initialize services
	athleteScoreService := services.NewAthleteScoreService(athleteScoreRepo)
	var athleteService interfaces.AthleteService = services.NewAthleteService(athleteRepo)
	var outcomeService interfaces.OutcomeService = services.NewOutcomeService(outcomeRepo, athleteScoreService, boutRepo)
	boutService := services.NewBoutService(boutRepo)
	feedService := services.NewFeedService(feedRepo)
	gymService := services.NewGymService(gymRepo)
	styleService := services.NewStyleService(styleRepo, athleteScoreService)

	// Initialize handlers
	athleteHandler := services.NewAthleteHandler(athleteService)
	outcomeHandler := services.NewOutcomeHandler(outcomeService)
	athleteScoreHandler := services.NewAthleteScoreHandler(athleteScoreService)
	boutHandler := services.NewBoutHandler(boutService)
	feedHandler := services.NewFeedHandler(feedService)
	gymHandler := services.NewGymHandler(gymService)
	styleHandler := services.NewStyleHandler(styleService)

	// Set handlers in router package
	router.SetAthleteHandler(athleteHandler)
	router.SetOutcomeHandler(outcomeHandler)
	router.SetAthleteScoreHandler(athleteScoreHandler)
	router.SetBoutHandler(boutHandler)
	router.SetFeedHandler(feedHandler)
	router.SetGymHandler(gymHandler)
	router.SetStyleHandler(styleHandler)

	// Create router with all routes configured
	r := router.CreateRouter()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
