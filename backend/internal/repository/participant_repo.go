package repository

import (
	"database/sql"
	"log"
	"strings"

	"github.com/nycrat/duck-hunt/backend/internal/types"
)

type ParticipantRepo struct {
	db *sql.DB
}

type ParticipantRepositoryInterface interface {
	GetAllParticipants() ([]types.Participant, bool)
	GetAllParticipantById(id int) (types.Participant, bool)
}

func NewParticipantRepo(db *sql.DB) *ParticipantRepo {
	return &ParticipantRepo{
		db: db,
	}
}

func (r *ParticipantRepo) GetAllParticipants() ([]types.Participant, bool) {
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

func (r *ParticipantRepo) GetParticipantById(id int) (types.Participant, bool) {
	var name string
	var score int

	err := r.db.QueryRow(`SELECT name, score FROM participants WHERE id = $1`, id).Scan(&name, &score)

	if err != nil {
		log.Println(err)
		return types.Participant{}, false
	}

	return types.Participant{Id: id, Name: name, Score: score}, true
}
