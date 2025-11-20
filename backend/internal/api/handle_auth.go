package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/nycrat/duck-hunt/backend/internal/repository"
)

type AuthHandler struct {
	r      *repository.AuthRepo
	jwtKey []byte
}

func NewAuthHandler(r *repository.AuthRepo, jwtKey []byte) *AuthHandler {
	return &AuthHandler{
		r:      r,
		jwtKey: jwtKey,
	}
}

func (h *AuthHandler) HandlePostAuth(w http.ResponseWriter, r *http.Request) {
	scheme, passcode, found := strings.Cut(r.Header.Get("Authorization"), " ")

	if !found || scheme != "Basic" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	id, ok := h.r.GetAuthorizedId(passcode)

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, ok := repository.GenerateJwtToken(id, h.jwtKey)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

func (h *AuthHandler) HandlePostSession(w http.ResponseWriter, r *http.Request) {
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
func (h *AuthHandler) HandlePostAuthAdmin(w http.ResponseWriter, r *http.Request) {
	admin := r.Context().Value("admin").(bool)
	if admin {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
