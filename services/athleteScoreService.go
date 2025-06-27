package services

import (
	"errors"
	"log"
	"math"
	"strconv"

	"ronin/models"
	"ronin/repositories"
)

type AthleteScoreService struct {
	repo *repositories.AthleteScoreRepository
}

func NewAthleteScoreService(repo *repositories.AthleteScoreRepository) *AthleteScoreService {
	return &AthleteScoreService{
		repo: repo,
	}
}

func (s *AthleteScoreService) GetAllAthleteScoresByAthleteId(athleteId int) ([]models.AthleteScore, error) {
	scores, err := s.repo.GetAllAthleteScoresByAthleteId(strconv.Itoa(athleteId))
	if err != nil {
		return nil, err
	}
	return scores, nil
}

func (s *AthleteScoreService) GetAthleteScoreByStyle(athleteId, styleId int) (models.AthleteScore, error) {
	score, err := s.repo.GetAthleteScoreByStyle(athleteId, styleId)
	if err != nil {
		return models.AthleteScore{}, err
	}
	return score, nil
}

func (s *AthleteScoreService) CreateAthleteScoreUponRegistration(athleteId, styleId int) error {
	return s.repo.CreateAthleteScoreUponRegistration(athleteId, styleId)
}

func (s *AthleteScoreService) CalculateNewScores(winnerScore, loserScore *models.AthleteScore, isDraw bool, outcomeId int) error {
	if winnerScore == nil || loserScore == nil {
		return errors.New("both winner and loser scores must be provided")
	}

	// Calculate Elo rating changes
	expectedScoreWinner := 1.0 / (1.0 + math.Pow(10, (loserScore.Score-winnerScore.Score)/400.0))
	expectedScoreLoser := 1.0 / (1.0 + math.Pow(10, (winnerScore.Score-loserScore.Score)/400.0))

	// K-factor for rating adjustment (can be adjusted based on requirements)
	k := 32.0

	var newWinnerScore, newLoserScore float64
	if isDraw {
		newWinnerScore = winnerScore.Score + k*(0.5-expectedScoreWinner)
		newLoserScore = loserScore.Score + k*(0.5-expectedScoreLoser)
	} else {
		newWinnerScore = winnerScore.Score + k*(1.0-expectedScoreWinner)
		newLoserScore = loserScore.Score + k*(0.0-expectedScoreLoser)
	}

	// Update scores in repository
	if err := s.repo.UpdateAthleteScore(int(newWinnerScore), winnerScore.AthleteId, winnerScore.StyleId, outcomeId); err != nil {
		log.Printf("Error updating winner score: %v", err)
		return err
	}

	if err := s.repo.UpdateAthleteScore(int(newLoserScore), loserScore.AthleteId, loserScore.StyleId, outcomeId); err != nil {
		log.Printf("Error updating loser score: %v", err)
		return err
	}

	return nil
}

func (s *AthleteScoreService) GetAthleteScoreHistoryByStyle(athleteId, styleId int) (models.AthleteScoreHistoryResponse, error) {
	history, styleName, err := s.repo.GetAthleteScoreHistoryByAthleteAndStyle(athleteId, styleId)
	if err != nil {
		return models.AthleteScoreHistoryResponse{}, err
	}

	// Convert to response format
	var historyEntries []models.ScoreHistoryEntry
	for _, entry := range history {
		historyEntries = append(historyEntries, models.ScoreHistoryEntry{
			Date:  entry.CreatedDate,
			Score: entry.NewScore,
		})
	}

	response := models.AthleteScoreHistoryResponse{
		StyleId:   styleId,
		StyleName: styleName,
		History:   historyEntries,
	}

	return response, nil
}
