package models

type AthleteScoreHistory struct {
	HistoryId     int    `json:"historyId" db:"history_id"`
	AthleteId     int    `json:"athleteId" db:"athlete_id"`
	StyleId       int    `json:"styleId" db:"style_id"`
	OutcomeId     int    `json:"outcomeId" db:"outcome_id"`
	PreviousScore int    `json:"previousScore" db:"previous_score"`
	NewScore      int    `json:"newScore" db:"new_score"`
	CreatedDate   string `json:"createdDate" db:"created_dt"`
	UpdatedDate   string `json:"updatedDate" db:"updated_dt"`
}

type AthleteScoreHistoryResponse struct {
	StyleId   int                           `json:"styleId"`
	StyleName string                        `json:"styleName"`
	History   []AthleteScoreHistoryEntry    `json:"history"`
}

type AthleteScoreHistoryEntry struct {
	Date  string `json:"date"`
	Score int    `json:"score"`
}