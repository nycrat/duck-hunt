package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/nycrat/duck-hunt/backend/internal/repository"
)

func HandlePostAuth(w http.ResponseWriter, r *http.Request) {
	scheme, passcode, found := strings.Cut(r.Header.Get("Authorization"), " ")

	if !found || scheme != "Basic" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	pepper := r.Context().Value("pepper").([]byte)
	db := r.Context().Value("db").(*sql.DB)

	id, ok := repository.DbSelectId(passcode, pepper, db)

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	hs256Key := r.Context().Value("key").([]byte)

	token, ok := repository.GenerateJwtToken(id, hs256Key)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

func HandlePostSession(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := json.Marshal(id)

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(body)
}

// Handles /auth/admin, returns OK status if user has session id = 1
func HandlePostAuthAdmin(w http.ResponseWriter, r *http.Request) {
	admin := r.Context().Value("admin").(bool)
	if admin {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
