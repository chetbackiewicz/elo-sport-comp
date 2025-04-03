package repositories

import (
	"log"
	"ronin/models"

	"github.com/jmoiron/sqlx"
)

type AthleteRepository struct {
	db *sqlx.DB
}

type AthleteUsername struct {
	Username string `json:"username" db:"username"`
}

func NewAthleteRepository(db *sqlx.DB) *AthleteRepository {
	return &AthleteRepository{
		db: db,
	}
}

func (repo *AthleteRepository) GetAllUsernames() ([]string, error) {
	var usernames []string
	var tempUsername AthleteUsername

	sqlStmt := `SELECT username FROM athlete`
	rows, err := repo.db.Queryx(sqlStmt)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.StructScan(&tempUsername)
		if err != nil {
			return nil, err
		}
		usernames = append(usernames, tempUsername.Username)
	}

	return usernames, nil
}

func (repo *AthleteRepository) GetAllAthletes() ([]models.Athlete, error) {
	var athletes []models.Athlete
	var tempAthlete models.Athlete

	sqlStmt := `SELECT * FROM athlete`
	rows, err := repo.db.Queryx(sqlStmt)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.StructScan(&tempAthlete)
		if err != nil {
			return nil, err
		}
		athletes = append(athletes, tempAthlete)
	}

	return athletes, nil
}

func (repo *AthleteRepository) GetAthleteById(id string) (models.Athlete, error) {
	var tempAthlete models.Athlete

	sqlStmt := `SELECT * FROM athlete where athlete_id = $1`
	err := repo.db.Get(&tempAthlete, sqlStmt, id)
	if err != nil {
		return tempAthlete, err
	}

	return tempAthlete, nil
}

func (repo *AthleteRepository) GetAthleteByUsername(username string) (models.Athlete, error) {
	var tempAthlete models.Athlete

	sqlStmt := `SELECT * FROM athlete where username = $1`
	err := repo.db.Get(&tempAthlete, sqlStmt, username)
	if err != nil {
		return tempAthlete, err
	}

	return tempAthlete, nil
}

func (repo *AthleteRepository) IsAuthorizedUser(athlete models.Athlete) (bool, models.Athlete, error) {
	log.Printf("Checking authorization for user: %+v", athlete)

	var athleteId int
	sqlStmt := `SELECT count(*) FROM athlete where username = $1 and password = $2`
	log.Printf("Executing SQL: %s with params: username=%s", sqlStmt, athlete.Username)

	err := repo.db.QueryRow(sqlStmt, athlete.Username, athlete.Password).Scan(&athleteId)
	if err != nil {
		log.Printf("Error in initial auth check: %v", err)
		return false, models.Athlete{}, err
	}
	log.Printf("Found %d matching users", athleteId)

	if athleteId == 1 {
		var tempAthlete models.Athlete
		sqlStmt := `SELECT * FROM athlete where username = $1 and password = $2`
		log.Printf("Fetching athlete details with SQL: %s", sqlStmt)

		err := repo.db.Get(&tempAthlete, sqlStmt, athlete.Username, athlete.Password)
		if err != nil {
			log.Printf("Error fetching athlete details: %v", err)
			return true, models.Athlete{}, err
		}
		log.Printf("Successfully retrieved athlete: %+v", tempAthlete)
		return true, tempAthlete, nil
	}

	log.Println("No matching user found")
	return false, models.Athlete{}, nil
}

func (repo *AthleteRepository) CreateAthlete(athlete models.Athlete) (int, error) {
	var athleteId int
	sqlStmt := `INSERT INTO athlete (first_name, last_name, username, birth_date, email, password)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING athlete_id`
	err := repo.db.QueryRow(sqlStmt, athlete.FirstName, athlete.LastName, athlete.Username, athlete.BirthDate, athlete.Email, athlete.Password).Scan(&athleteId)
	if err != nil {
		return 0, err
	}

	sqlStmt = `INSERT INTO athlete_record (athlete_id, wins, losses, draws) VALUES ($1, 0, 0, 0)`
	_, err = repo.db.Exec(sqlStmt, athleteId)
	if err != nil {
		return 0, err
	}

	return athleteId, nil
}

func (repo *AthleteRepository) UpdateAthlete(athlete models.Athlete) error {
	sqlStmt := `UPDATE athlete SET first_name = $1, last_name = $2, username = $3, birth_date = $4, email = $5, password = $6 WHERE athlete_id = $7`
	_, err := repo.db.Exec(sqlStmt, athlete.FirstName, athlete.LastName, athlete.Username, athlete.BirthDate, athlete.Email, athlete.Password, athlete.AthleteId)
	return err
}

func (repo *AthleteRepository) DeleteAthlete(id string) error {
	sqlStmt := `DELETE FROM athlete WHERE athlete_id = $1`
	_, err := repo.db.Exec(sqlStmt, id)
	return err
}

func (repo *AthleteRepository) GetAthleteRecord(id string) (models.AthleteRecord, error) {
	var record models.AthleteRecord
	sqlStmt := `SELECT * FROM athlete_record where athlete_id = $1`
	err := repo.db.Get(&record, sqlStmt, id)
	if err != nil {
		return record, err
	}
	return record, nil
}

func (repo *AthleteRepository) FollowAthlete(follow models.Follow) error {
	sqlStmt := `INSERT INTO following (follower_id, followed_id) VALUES ($1, $2)`
	_, err := repo.db.Exec(sqlStmt, follow.FollowerId, follow.FollowedId)
	return err
}

func (repo *AthleteRepository) UnfollowAthlete(followerId, followedId int) error {
	sqlStmt := `DELETE FROM following WHERE follower_id = $1 AND followed_id = $2`
	_, err := repo.db.Exec(sqlStmt, followerId, followedId)
	return err
}

func (repo *AthleteRepository) GetAthletesFollowed(id string) ([]int, error) {
	var follows []int
	var tempFollow models.Follow
	sqlStmt := `SELECT * FROM following where follower_id = $1`
	rows, err := repo.db.Queryx(sqlStmt, id)
	if err != nil {
		return follows, err
	}
	defer rows.Close()

	for rows.Next() {
		err2 := rows.StructScan(&tempFollow)
		if err2 != nil {
			return follows, err2
		}
		follows = append(follows, tempFollow.FollowedId)
	}

	return follows, nil
}
