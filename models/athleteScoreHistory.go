package models

import "time"

// AthleteScoreHistory represents individual score history entries
type ScoreHistoryEntry struct {
	Date  time.Time `json:"date" db:"created_dt"`
	Score int       `json:"score" db:"new_score"`
}

// AthleteScoreHistory represents the response format for the history endpoint
type AthleteScoreHistory struct {
	StyleId   int                 `json:"styleId" db:"style_id"`
	StyleName string              `json:"styleName" db:"style_name"`
	History   []ScoreHistoryEntry `json:"history"`
}
