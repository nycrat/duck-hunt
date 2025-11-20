package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/nycrat/duck-hunt/backend/internal/repository"
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

	submissions, ok := repository.DbFetchSubmissions(db, id, title)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	serialized, _ := json.Marshal(submissions)
	w.Write(serialized)
}

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
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	repository.DbPostNewSubmission(db, id.(int), title, image)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

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

	repository.DbPostReview(db, int(submissionId), string(status))
}
