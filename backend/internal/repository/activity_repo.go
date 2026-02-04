package repository

import (
	"database/sql"
	"log"
	"strings"

	"github.com/nycrat/duck-hunt/backend/internal/types"
)

type ActivityRepository struct {
	db *sql.DB
}

type ActivityRepositoryInterface interface {
	GetActivityList(db *sql.DB) ([]types.Activity, bool)
	GetActivityByTitle(db *sql.DB, title string) (types.Activity, bool)
}

func NewActivityRepository(db *sql.DB) *ActivityRepository {
	return &ActivityRepository{
		db: db,
	}
}

func (r *ActivityRepository) GetActivityList() ([]types.Activity, bool) {
	rows, err := r.db.Query("SELECT title, points, description FROM activities ORDER BY title")

	if err != nil {
		log.Println(err)
		return []types.Activity{}, false
	}

	activities := []types.Activity{}

	for rows.Next() {
		var title string
		var points int
		var description string
		err := rows.Scan(&title, &points, &description)
		if err != nil {
			log.Println(err)
			return []types.Activity{}, false
		}

		title = strings.TrimSpace(title)
		activities = append(activities, types.Activity{Title: title, Points: points, Description: description})
	}

	return activities, true
}

func (r *ActivityRepository) GetActivityByTitle(title string) (types.Activity, bool) {
	var points int
	var description string
	err := r.db.QueryRow(`SELECT points, description FROM activities WHERE title = $1`, title).Scan(&points, &description)

	if err != nil {
		return types.Activity{}, false
	}

	return types.Activity{
		Title:       title,
		Points:      points,
		Description: description,
	}, true
}
