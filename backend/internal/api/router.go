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

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", authHandler.HandlePostAuth)
		r.Post("/session", authHandler.HandlePostSession)
		r.Post("/admin", authHandler.HandlePostAuthAdmin)
	})

	r.Route("/participants", func(r chi.Router) {
		r.Get("/", participantHandler.HandleGetParticipantList)
		r.Get("/{id}", participantHandler.HandleGetParticipantInfo)
	})

	r.Route("/activities", func(r chi.Router) {
		r.Get("/", activityHandler.HandleGetActivityList)
		r.Get("/{title}", activityHandler.HandleGetActivity)

		r.Route("/{title}/submissions", func(r chi.Router) {
			r.Get("/", submissionHandler.HandleGetSubmissionList)
			r.Post("/", submissionHandler.HandlePostSubmission)
		})
	})

	r.Route("/admin", func(r chi.Router) {
		r.Route("/submissions", func(r chi.Router) {
			r.Get("/", submissionHandler.HandleGetUnreviewedSubmissionList)
			r.Patch("/{id}", submissionHandler.HandleUpdateSubmission)
		})
	})

	return r
}
