package repository

import (
	"database/sql"
	"log"
	"strings"

	"github.com/nycrat/duck-hunt/backend/internal/types"
)

func DbFetchParticipants(db *sql.DB) ([]types.Participant, bool) {
	rows, err := db.Query("SELECT id, name, score FROM participants ORDER BY id")

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

func DbFetchParticipantById(db *sql.DB, id int) (types.Participant, bool) {
	var name string
	var score int

	err := db.QueryRow(`SELECT name, score FROM participants WHERE id = $1`, id).Scan(&name, &score)

	if err != nil {
		log.Println(err)
		return types.Participant{}, false
	}

	return types.Participant{Id: id, Name: name, Score: score}, true
}
