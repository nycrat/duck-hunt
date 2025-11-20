package repository

import (
	"database/sql"
	"log"

	"github.com/nycrat/duck-hunt/backend/internal/types"
)

type SubmissionRepository struct {
	db *sql.DB
}

type SubmissionRepositoryInterface interface {
	GetAllUserSubmissionsForActivity(db *sql.DB, id int, title string) ([]types.Submission, bool)
	AddNewSubmission(db *sql.DB, id int, title string, image []byte)
	GetNumberOfUserSubmissionsForActivity(db *sql.DB, id int, title string) (int, bool)
	UpdateSubmissionStatus(db *sql.DB, submissionId int, status string)
}

func NewSubmissionRepository(db *sql.DB) *SubmissionRepository {
	return &SubmissionRepository{
		db: db,
	}
}

func (r *SubmissionRepository) GetAllUserSubmissionsForActivity(id int, title string) ([]types.Submission, bool) {
	rows, err := r.db.Query(`SELECT id, status, image FROM submissions WHERE participant_id = $1 AND activity_title = $2 ORDER BY id`, id, title)

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

func (r *SubmissionRepository) AddNewSubmission(id int, title string, image []byte) {
	_, err := r.db.Query(`INSERT INTO submissions (participant_id, activity_title, image) VALUES($1, $2, $3)`, id, title, image)

	if err != nil {
		log.Println(err)
	}
}

func (r *SubmissionRepository) GetNumberOfUserSubmissionsForActivity(id int, title string) (int, bool) {
	var count int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM submissions WHERE participant_id = $1 AND activity_title = $2`, id, title).Scan(&count)

	if err != nil {
		log.Println(err)
		return 0, false
	}

	return count, true
}

func (r *SubmissionRepository) UpdateSubmissionStatus(submissionId int, status string) {
	_, err := r.db.Query(`UPDATE submissions SET status = $1 WHERE id = $2`, status, submissionId)

	if err != nil {
		log.Println(err)
	}
}
