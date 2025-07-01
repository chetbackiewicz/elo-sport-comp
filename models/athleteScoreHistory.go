package models

type AthleteScoreHistoryEntry struct {
	Date  string `json:"date" db:"created_dt"`
	Score int    `json:"score" db:"new_score"`
}

type AthleteScoreHistory struct {
	StyleId   int                        `json:"styleId" db:"style_id"`
	StyleName string                     `json:"styleName" db:"style_name"`
	History   []AthleteScoreHistoryEntry `json:"history"`
}