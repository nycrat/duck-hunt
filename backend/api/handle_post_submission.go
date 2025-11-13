package api

import (
	"backend/common"
	"database/sql"
	"io"
	"log"
	"net/http"
)

func HandlePostSubmission(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	db := r.Context().Value("db").(*sql.DB)

	title := r.PathValue("title")

	image, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	common.DbPostNewSubmission(db, id.(int), title, image)

	if err != nil {
		log.Fatal(err)
	}
}
