package repositories

import (
	"fmt"
	"ronin/models"

	"github.com/jmoiron/sqlx"
)

type FeedRepository struct {
	DB *sqlx.DB
}

func NewFeedRepository(db *sqlx.DB) *FeedRepository {
	return &FeedRepository{
		DB: db,
	}
}

func (fr *FeedRepository) GetFeedByAthleteId(id string) ([]models.Feed, error) {
	var feed []models.Feed

	sqlStmt := `WITH athlete_following AS (
		SELECT followed_id
		FROM following
		WHERE follower_id = $1
	),
	latest_scores AS (
		SELECT athlete_id, style_id, score, updated_dt,
			   ROW_NUMBER() OVER (PARTITION BY athlete_id, style_id ORDER BY updated_dt DESC) as row_num
		FROM athlete_score
	)
	SELECT
		DISTINCT b.bout_id AS "boutId",
		c.first_name AS "challengerFirstName",
		c.last_name AS "challengerLastName",
		c.username AS "challengerUsername",
		c.athlete_id AS "challengerId",
		a.first_name AS "acceptorFirstName",
		a.athlete_id AS "acceptorId",
		a.last_name AS "acceptorLastName",
		a.username AS "acceptorUsername",
		w.first_name AS "winnerFirstName",
		w.last_name AS "winnerLastName",
		w.username AS "winnerUsername",
		l.first_name AS "loserFirstName",
		l.last_name AS "loserLastName",
		l.username AS "loserUsername",
		o.is_draw AS "isDraw",
		r.first_name AS "refereeFirstName",
		r.last_name AS "refereeLastName",
		r.athlete_id AS "refereeId",
		s.style_id AS "styleId",
		s.style_name AS "style",
		o.winner_id AS "winnerId",
		o.loser_id AS "loserId",
		b.updated_dt AS "updatedDt",
		COALESCE(ww.wins, 0) AS "winnerWins",
		COALESCE(ww.losses, 0) AS "winnerLosses",
		COALESCE(ww.draws, 0) AS "winnerDraws",
		COALESCE(ll.wins, 0) AS "loserWins",
		COALESCE(ll.losses, 0) AS "loserLosses",
		COALESCE(ll.draws, 0) AS "loserDraws",
		COALESCE(ws.score, 0) AS "winnerScore",
		COALESCE(ls.score, 0) AS "loserScore"
	FROM
		bout b
	JOIN
		athlete c ON b.challenger_id = c.athlete_id
	JOIN
		athlete a ON b.acceptor_id = a.athlete_id
	JOIN
		outcome o ON b.bout_id = o.bout_id
	JOIN
		athlete w ON o.winner_id = w.athlete_id
	JOIN
		athlete l ON o.loser_id = l.athlete_id
	JOIN
		athlete r ON b.referee_id = r.athlete_id
	JOIN
		style s ON b.style_id = s.style_id
	LEFT JOIN
		athlete_record ww ON o.winner_id = ww.athlete_id
	LEFT JOIN
		athlete_record ll ON o.loser_id = ll.athlete_id
	LEFT JOIN
		latest_scores ws ON o.winner_id = ws.athlete_id AND ws.style_id = b.style_id AND ws.row_num = 1
	LEFT JOIN
		latest_scores ls ON o.loser_id = ls.athlete_id AND ls.style_id = b.style_id AND ls.row_num = 1
	WHERE 
		b.cancelled != true 
		AND b.completed = true 
		AND b.accepted = true
		AND (
			-- Include bouts where the user is a participant
			b.challenger_id = $1 
			OR b.acceptor_id = $1
			-- Include bouts where someone they follow is a participant
			OR (b.challenger_id IN (SELECT followed_id FROM athlete_following)
				OR b.acceptor_id IN (SELECT followed_id FROM athlete_following))
		)
	ORDER BY b.updated_dt DESC;`

	rows, err := fr.DB.Queryx(sqlStmt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	feed = []models.Feed{} // Initialize as empty slice instead of nil
	var tempFeed models.Feed

	for rows.Next() {
		err = rows.StructScan(&tempFeed)
		if err != nil {
			// Log the error but continue with other rows
			fmt.Printf("Error scanning row: %v\n", err)
			continue
		}
		feed = append(feed, tempFeed)
	}

	// Check for errors after iterating through rows
	if err = rows.Err(); err != nil {
		return feed, err
	}

	// If we got no results, still return an empty slice instead of nil
	if len(feed) == 0 {
		return []models.Feed{}, nil
	}

	return feed, nil
}
