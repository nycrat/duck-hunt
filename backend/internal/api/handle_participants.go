package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/nycrat/duck-hunt/backend/internal/repository"
	"github.com/nycrat/duck-hunt/backend/internal/types"
)

func HandleGetParticipants(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	db := r.Context().Value("db").(*sql.DB)

	participants, ok := repository.DbFetchParticipants(db)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	serialized, _ := json.Marshal(participants)
	w.Write(serialized)
}

func HandleGetParticipantSubmissionCounts(w http.ResponseWriter, r *http.Request) {
	admin := r.Context().Value("admin").(bool)

	if !admin {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	db := r.Context().Value("db").(*sql.DB)
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 32)

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	activities, ok := repository.DbFetchActivityPreviews(db)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := []types.ActivitySubmissions{}

	for _, activity := range activities {
		count, ok := repository.DbCountNumberOfSubmissions(db, int(id), activity.Title)

		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res = append(res, types.ActivitySubmissions{
			Title: activity.Title,
			Count: count,
		})
	}

	serialized, _ := json.Marshal(res)

	w.Write(serialized)
}

func HandleGetParticipantInfo(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	db := r.Context().Value("db").(*sql.DB)
	targetId, err := strconv.ParseInt(r.PathValue("id"), 10, 32)

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	participants, ok := repository.DbFetchParticipantById(db, int(targetId))

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	serialized, _ := json.Marshal(participants)
	w.Write(serialized)
}
