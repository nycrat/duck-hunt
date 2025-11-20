package api

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/nycrat/duck-hunt/backend/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func DuckHuntRouter(jwtKey []byte, pepper []byte, db *sql.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*"},

		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	participantHandler := NewParticipantHandler(repository.NewParticipantRepo(db))
	authHandler := NewAuthHandler(repository.NewAuthRepo(db, pepper), jwtKey)
	activityHandler := NewActivityHandler(repository.NewActivityRepository(db))
	submissionHandler := NewSubmissionHandler(repository.NewSubmissionRepository(db), repository.NewActivityRepository(db))

	r.Use(GetJwtMiddleware(jwtKey))

	r.Post("/auth", authHandler.HandlePostAuth)
	r.Post("/auth/admin", authHandler.HandlePostAuthAdmin)
	r.Post("/session", authHandler.HandlePostSession)

	r.Get("/participants", participantHandler.HandleGetParticipants)
	r.Get("/participants/{id}", participantHandler.HandleGetParticipantInfo)

	r.Get("/activities", activityHandler.HandleGetActivityPreviews)
	r.Get("/activities/{title}", activityHandler.HandleGetActivity)

	r.Get("/submissions/{title}", submissionHandler.HandleGetSubmissions)
	r.Get("/submissions/{title}/{id}", submissionHandler.HandleGetSubmissions)

	r.Post("/submissions/{title}", submissionHandler.HandlePostSubmission)

	r.Get("/participants/{id}/submission_counts", submissionHandler.HandleGetParticipantSubmissionCounts)

	r.Post("/review/{submissionId}", submissionHandler.HandlePostReview)

	return r
}
