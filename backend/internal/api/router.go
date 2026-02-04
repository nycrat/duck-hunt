package api

import (
	"database/sql"
	"net/http"

	"github.com/nycrat/duck-hunt/backend/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func DuckHuntRouter(jwtKey []byte, pepper []byte, db *sql.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},

		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	participantHandler := NewParticipantHandler(repository.NewParticipantRepository(db))
	authHandler := NewAuthHandler(repository.NewAuthRepo(db, pepper), jwtKey)
	activityHandler := NewActivityHandler(repository.NewActivityRepository(db))
	submissionHandler := NewSubmissionHandler(repository.NewSubmissionRepository(db), repository.NewActivityRepository(db), repository.NewParticipantRepository(db))

	r.Use(GetJwtMiddleware(jwtKey))

	r.Route("/participants", func(r chi.Router) {
		r.Get("/", participantHandler.HandleGetParticipants)
		r.Get("/{id}", participantHandler.HandleGetParticipantInfo)
		r.Get("/{id}/submission_counts", submissionHandler.HandleGetParticipantSubmissionCounts)
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", authHandler.HandlePostAuth)
		r.Post("/admin", authHandler.HandlePostAuthAdmin)
		r.Post("/session", authHandler.HandlePostSession)
	})

	r.Route("/activities", func(r chi.Router) {
		r.Get("/", activityHandler.HandleGetActivityPreviews)
		r.Get("/{title}", activityHandler.HandleGetActivity)
	})

	r.Route("/submissions", func(r chi.Router) {
		r.Get("/list/unreviewed/todo", submissionHandler.HandleGetUnreviewedSubmissions)

		r.Get("/{title}", submissionHandler.HandleGetSubmissions)
		r.Get("/{title}/{id}", submissionHandler.HandleGetSubmissions)
		r.Post("/{title}", submissionHandler.HandlePostSubmission)

		r.Post("/review/{submissionId}", submissionHandler.HandlePostReview)
	})

	return r
}
