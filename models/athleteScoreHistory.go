package models

import (
	"time"
)

type AthleteScoreHistory struct {
	HistoryId     int       `json:"historyId" db:"history_id"`
	AthleteId     int       `json:"athleteId" db:"athlete_id"`
	StyleId       int       `json:"styleId" db:"style_id"`
	OutcomeId     *int      `json:"outcomeId,omitempty" db:"outcome_id"`
	PreviousScore *int      `json:"previousScore,omitempty" db:"previous_score"`
	NewScore      int       `json:"score" db:"new_score"`
	CreatedDate   time.Time `json:"date" db:"created_dt"`
	UpdatedDate   time.Time `json:"updatedDate" db:"updated_dt"`
}

// Represents the response for the style score history
type AthleteStyleScoreHistory struct {
	StyleId   int                  `json:"styleId"`
	StyleName string               `json:"styleName"`
	History   []ScoreHistoryEntry  `json:"history"`
}

// Individual score history entry for the response
type ScoreHistoryEntry struct {
	Date  time.Time `json:"date" db:"created_dt"`
	Score int       `json:"score" db:"new_score"`
}