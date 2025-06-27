package models

import "time"

type AthleteScoreHistory struct {
	HistoryId     int       `json:"historyId" db:"history_id"`
	AthleteId     int       `json:"athleteId" db:"athlete_id"`
	StyleId       int       `json:"styleId" db:"style_id"`
	OutcomeId     *int      `json:"outcomeId" db:"outcome_id"`
	PreviousScore *int      `json:"previousScore" db:"previous_score"`
	NewScore      int       `json:"newScore" db:"new_score"`
	CreatedDate   time.Time `json:"createdDate" db:"created_dt"`
	UpdatedDate   time.Time `json:"updatedDate" db:"updated_dt"`
}

type ScoreHistoryEntry struct {
	Date  time.Time `json:"date"`
	Score int       `json:"score"`
}

type AthleteScoreHistoryResponse struct {
	StyleId   int                 `json:"styleId"`
	StyleName string              `json:"styleName"`
	History   []ScoreHistoryEntry `json:"history"`
}