package repository

import (
	"database/sql"
	"log"
	"strings"

	"github.com/nycrat/duck-hunt/backend/internal/types"
)

func DbFetchActivityPreviews(db *sql.DB) ([]types.ActivityPreview, bool) {
	rows, err := db.Query("SELECT title, points FROM activities ORDER BY title")

	if err != nil {
		log.Println(err)
		return []types.ActivityPreview{}, false
	}

	activities := []types.ActivityPreview{}

	for rows.Next() {
		var title string
		var points int
		err := rows.Scan(&title, &points)
		if err != nil {
			log.Println(err)
			return []types.ActivityPreview{}, false
		}

		title = strings.TrimSpace(title)
		activities = append(activities, types.ActivityPreview{Title: title, Points: points})
	}

	return activities, true
}

func DbFetchActivity(db *sql.DB, title string) (types.Activity, bool) {
	var points int
	var description string
	err := db.QueryRow(`SELECT points, description FROM activities WHERE title = $1`, title).Scan(&points, &description)

	if err != nil {
		return types.Activity{}, false
	}

	return types.Activity{
		Title:       title,
		Points:      points,
		Description: description,
	}, true
}
