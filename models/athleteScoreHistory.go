package models

import "time"

type AthleteScoreHistory struct {
	HistoryID     int       `db:"history_id"`
	AthleteID     int       `db:"athlete_id"`
	StyleID       int       `db:"style_id"`
	OutcomeID     *int      `db:"outcome_id"`
	PreviousScore *int      `db:"previous_score"`
	NewScore      int       `db:"new_score"`
	CreatedDt     time.Time `db:"created_dt"`
	UpdatedDt     time.Time `db:"updated_dt"`
}
