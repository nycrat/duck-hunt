package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/nycrat/duck-hunt/backend/internal/repository"
)

type SubmissionPatch struct {
	Id            int    `json:"id"`
	Status        string `json:"status"`
	ParticipantId int    `json:"participant_id"`
}

type SubmissionHandler struct {
	s *repository.SubmissionRepository
	a *repository.ActivityRepository
	p *repository.ParticipantRepository
}

func NewSubmissionHandler(s *repository.SubmissionRepository, a *repository.ActivityRepository, p *repository.ParticipantRepository) *SubmissionHandler {
	return &SubmissionHandler{
		s: s,
		a: a,
		p: p,
	}
}

func (h *SubmissionHandler) HandleGetSubmissionList(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	title := r.PathValue("title")

	submissions, ok := h.s.GetAllUserSubmissionsForActivity(id.(int), title)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	serialized, _ := json.Marshal(submissions)
	w.Write(serialized)
}

func (h *SubmissionHandler) HandleGetUnreviewedSubmissionList(w http.ResponseWriter, r *http.Request) {
	admin := r.Context().Value("admin").(bool)
	if !admin {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	submissions, ok := h.s.GetUnreviewedSubmissions()

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	serialized, _ := json.Marshal(submissions)
	w.Write(serialized)
}

func (h *SubmissionHandler) HandlePostSubmission(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")
	if id == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	title := r.PathValue("title")

	image, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	h.s.AddNewSubmission(id.(int), title, image)
}

func (h *SubmissionHandler) HandleUpdateSubmission(w http.ResponseWriter, r *http.Request) {
	admin := r.Context().Value("admin").(bool)
	if !admin {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var patch SubmissionPatch

	err = json.Unmarshal(body, &patch)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.s.UpdateSubmissionStatus(patch.Id, patch.Status)
	h.p.UpdateParticipantScore(patch.ParticipantId)
}
