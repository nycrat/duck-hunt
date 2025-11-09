package api

import (
	"backend/common"
	"database/sql"
	"encoding/json"
	"net/http"
)

func HandleGetParticipants(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	db := r.Context().Value("db").(*sql.DB)

	participants := common.DbFetchParticipants(db)
	serialized, _ := json.Marshal(participants)
	w.Write(serialized)
}
