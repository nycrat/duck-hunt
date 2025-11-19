package api

import "net/http"

// Handles /auth/admin, returns OK status if user has session id = 1
func HandlePostAuthAdmin(w http.ResponseWriter, r *http.Request) {
	admin := r.Context().Value("admin").(bool)
	if admin {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
