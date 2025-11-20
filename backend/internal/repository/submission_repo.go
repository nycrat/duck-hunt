package repository

import (
	"database/sql"
	"log"

	"github.com/nycrat/duck-hunt/backend/internal/types"
)

func DbFetchSubmissions(db *sql.DB, id int, title string) ([]types.Submission, bool) {
	rows, err := db.Query(`SELECT id, status, image FROM submissions WHERE participant_id = $1 AND activity_title = $2 ORDER BY id`, id, title)

	if err != nil {
		log.Println(err)
		return []types.Submission{}, false
	}

	submissions := []types.Submission{}

	for rows.Next() {
		var submissionId int
		var status string
		var image []byte
		err := rows.Scan(&submissionId, &status, &image)
		if err != nil {
			log.Println(err)
			return []types.Submission{}, false
		}

		submissions = append(submissions, types.Submission{Id: submissionId, Status: status, Image: image})
	}

	return submissions, true
}

func DbPostNewSubmission(db *sql.DB, id int, title string, image []byte) {
	_, err := db.Query(`INSERT INTO submissions (participant_id, activity_title, image) VALUES($1, $2, $3)`, id, title, image)

	if err != nil {
		log.Println(err)
	}
}

func DbCountNumberOfSubmissions(db *sql.DB, id int, title string) (int, bool) {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM submissions WHERE participant_id = $1 AND activity_title = $2`, id, title).Scan(&count)

	if err != nil {
		log.Println(err)
		return 0, false
	}

	return count, true
}

func DbPostReview(db *sql.DB, submissionId int, status string) {
	_, err := db.Query(`UPDATE submissions SET status = $1 WHERE id = $2`, status, submissionId)

	if err != nil {
		log.Println(err)
	}
}
