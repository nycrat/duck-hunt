package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/nycrat/duck-hunt/backend/internal/repository"
)

type ParticipantHandler struct {
	r *repository.ParticipantRepository
}

func NewParticipantHandler(r *repository.ParticipantRepository) *ParticipantHandler {
	return &ParticipantHandler{
		r: r,
	}
}

func (h *ParticipantHandler) HandleGetParticipants(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	participants, ok := h.r.GetAllParticipants()

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	serialized, _ := json.Marshal(participants)
	w.Write(serialized)
}

func (h *ParticipantHandler) HandleGetParticipantInfo(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	targetId, err := strconv.ParseInt(r.PathValue("id"), 10, 32)

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	participants, ok := h.r.GetParticipantById(int(targetId))

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	serialized, _ := json.Marshal(participants)
	w.Write(serialized)
}
