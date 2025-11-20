package api

import (
	"encoding/json"
	"net/http"

	"github.com/nycrat/duck-hunt/backend/internal/repository"
)

type ActivityHandler struct {
	r *repository.ActivityRepository
}

func NewActivityHandler(r *repository.ActivityRepository) *ActivityHandler {
	return &ActivityHandler{
		r: r,
	}
}

func (h *ActivityHandler) HandleGetActivity(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	title := r.PathValue("title")

	activity, ok := h.r.GetActivityByTitle(title)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialized, _ := json.Marshal(activity)
	w.Write(serialized)
}

func (h *ActivityHandler) HandleGetActivityPreviews(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	activities, ok := h.r.GetAllActivityPreviews()

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	serialized, _ := json.Marshal(activities)
	w.Write(serialized)
}
