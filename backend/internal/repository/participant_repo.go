package repository

import (
	"database/sql"
	"log"
	"strings"

	"github.com/nycrat/duck-hunt/backend/internal/types"
)

type ParticipantRepository struct {
	db *sql.DB
}

type ParticipantRepositoryInterface interface {
	GetAllParticipants() ([]types.Participant, bool)
	GetAllParticipantById(id int) (types.Participant, bool)
	AddNewParticipant(name string)
}

func NewParticipantRepository(db *sql.DB) *ParticipantRepository {
	return &ParticipantRepository{
		db: db,
	}
}

func (r *ParticipantRepository) GetAllParticipants() ([]types.Participant, bool) {
	rows, err := r.db.Query("SELECT id, name, score FROM participants ORDER BY id")

	if err != nil {
		log.Println(err)
		return []types.Participant{}, false
	}

	participants := []types.Participant{}

	for rows.Next() {
		var id int
		var name string
		var score int
		err := rows.Scan(&id, &name, &score)
		if err != nil {
			log.Println(err)
			return []types.Participant{}, false
		}

		name = strings.TrimSpace(name)
		participants = append(participants, types.Participant{Id: id, Name: name, Score: score})
	}

	return participants, true
}

func (r *ParticipantRepository) GetParticipantById(id int) (types.Participant, bool) {
	var name string
	var score int

	err := r.db.QueryRow(`SELECT name, score FROM participants WHERE id = $1`, id).Scan(&name, &score)

	if err != nil {
		log.Println(err)
		return types.Participant{}, false
	}

	return types.Participant{Id: id, Name: name, Score: score}, true
}

func (r *ParticipantRepository) UpdateParticipantScore(id int) {
	_, err := r.db.Query(`
	UPDATE participants
	SET score = (
		SELECT COALESCE(SUM(points), 0) as total_score
		FROM activities
		WHERE title IN (
			SELECT activity_title
			FROM submissions
			WHERE participant_id = $1 AND status = 'accepted'
		)
	)
	WHERE participants.id = $1
	`, id)

	if err != nil {
		log.Println(err)
	}
}

func (r *ParticipantRepository) AddNewParticipant(name string) {
	_, err := r.db.Query(`INSERT INTO participants (name) VALUES ($1)`, name)

	if err != nil {
		log.Println(err)
	}
}
