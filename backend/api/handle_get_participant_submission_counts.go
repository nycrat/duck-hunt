package api

import (
	"backend/common"
	"backend/types"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

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

	activities, ok := common.DbFetchActivityPreviews(db)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := []types.ActivitySubmissions{}

	for _, activity := range activities {
		count, ok := common.DbCountNumberOfSubmissions(db, int(id), activity.Title)

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
