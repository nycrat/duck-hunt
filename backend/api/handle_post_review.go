package api

import (
	"backend/common"
	"database/sql"
	"io"
	"log"
	"net/http"
	"strconv"
)

func HandlePostReview(w http.ResponseWriter, r *http.Request) {
	admin := r.Context().Value("admin").(bool)
	if !admin {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	submissionId, err := strconv.ParseInt(r.PathValue("submissionId"), 10, 32)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	status, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	db := r.Context().Value("db").(*sql.DB)

	common.DbPostReview(db, int(submissionId), string(status))
}
