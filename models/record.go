package models

// Record represents an athlete's win/loss record
type Record struct {
	Wins   int `json:"wins" db:"wins"`
	Losses int `json:"losses" db:"losses"`
	Draws  int `json:"draws" db:"draws"`
}
