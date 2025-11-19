package api

import (
	"encoding/json"
	"log"
	"net/http"
)

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
