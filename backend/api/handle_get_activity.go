package api

import (
	"backend/common"
	"database/sql"
	"encoding/json"
	"net/http"
)

func HandleGetActivity(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	db := r.Context().Value("db").(*sql.DB)

	title := r.PathValue("title")

	activity, ok := common.DbFetchActivity(db, title)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialized, _ := json.Marshal(activity)
	w.Write(serialized)
}
