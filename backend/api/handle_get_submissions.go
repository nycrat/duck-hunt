package api

import (
	"backend/common"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func HandleGetSubmissions(w http.ResponseWriter, r *http.Request) {
	var id int

	pathId := r.PathValue("id")
	admin := r.Context().Value("admin").(bool)

	// Admin fetching a participant's submissions
	if pathId != "" {
		if !admin {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		parsedId, err := strconv.ParseInt(pathId, 10, 32)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		id = int(parsedId)
	} else {
		tokenId := r.Context().Value("id")
		if tokenId == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		id = tokenId.(int)
	}

	db := r.Context().Value("db").(*sql.DB)

	title := r.PathValue("title")

	submissions, ok := common.DbFetchSubmissions(db, id, title)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	serialized, _ := json.Marshal(submissions)
	w.Write(serialized)
}
