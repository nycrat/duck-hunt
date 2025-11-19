package api

import (
	"backend/common"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func HandleGetParticipantInfo(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	db := r.Context().Value("db").(*sql.DB)
	targetId, err := strconv.ParseInt(r.PathValue("id"), 10, 32)

	if err != nil {
		log.Fatal(err)
	}

	participants := common.DbFetchParticipantById(db, int(targetId))
	serialized, _ := json.Marshal(participants)
	w.Write(serialized)
}
