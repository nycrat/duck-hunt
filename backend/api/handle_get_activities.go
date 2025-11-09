package api

import (
	"backend/common"
	"database/sql"
	"encoding/json"
	"net/http"
)

func HandleGetActivites(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	db := r.Context().Value("db").(*sql.DB)

	activities := common.DbFetchActivities(db)
	serialized, _ := json.Marshal(activities)
	w.Write(serialized)
}
