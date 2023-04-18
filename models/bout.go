package models

type Bout struct {
	BoutId 			int `json:"boutId" db:"bout_id"`
	ChallengerId    int `json:"challengerId" db:"challenger_id"`
	AcceptorId 		int `json:"acceptorId" db:"acceptor_id"`
	Accepted  		bool `json:"accepted" db:"accepted"`
	Points  		int `json:"points" db:"points"`
}

func GetBout() Bout {
	var bout Bout
	return bout
}

func GetBouts() []Bout {
	var bouts []Bout
	return bouts
}

func CreateBout() Bout {
	var bout Bout
	return bout
}

func UpdateBout() Bout {
	var bout Bout
	return bout
}

func DeleteBout() Bout {
	var bout Bout
	return bout
}	
