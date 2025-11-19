package api

import (
	"backend/common"
	"database/sql"
	"encoding/json"
	"net/http"
)

func HandleGetSubmissions(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	db := r.Context().Value("db").(*sql.DB)

	title := r.PathValue("title")

	submissions, ok := common.DbFetchSubmissions(db, id.(int), title)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	serialized, _ := json.Marshal(submissions)
	w.Write(serialized)
}
