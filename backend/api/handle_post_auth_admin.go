package api

import "net/http"

// Handles /auth/admin, returns OK status if user has session id = 1
func HandlePostAuthAdmin(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == 1 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
