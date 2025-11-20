package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/nycrat/duck-hunt/backend/internal/repository"
)

func HandleGetActivity(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	db := r.Context().Value("db").(*sql.DB)

	title := r.PathValue("title")

	activity, ok := repository.DbFetchActivity(db, title)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialized, _ := json.Marshal(activity)
	w.Write(serialized)
}

func HandleGetActivityPreviews(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	db := r.Context().Value("db").(*sql.DB)

	activities, ok := repository.DbFetchActivityPreviews(db)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	serialized, _ := json.Marshal(activities)
	w.Write(serialized)
}
