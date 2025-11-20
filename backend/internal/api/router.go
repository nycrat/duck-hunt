package api

import (
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"

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

	r.Use(GetJwtMiddleware(jwtKey))
	r.Use(GetPepperMiddleware(pepper))
	r.Use(GetKeyMiddleware(jwtKey))
	r.Use(GetDbMiddleware(db))

	r.Post("/auth", HandlePostAuth)
	r.Post("/auth/admin", HandlePostAuthAdmin)
	r.Post("/session", HandlePostSession)

	r.Get("/participants", HandleGetParticipants)
	r.Get("/participants/{id}", HandleGetParticipantInfo)
	r.Get("/activities", HandleGetActivityPreviews)
	r.Get("/activities/{title}", HandleGetActivity)

	r.Get("/submissions/{title}", HandleGetSubmissions)
	r.Get("/submissions/{title}/{id}", HandleGetSubmissions)

	r.Post("/submissions/{title}", HandlePostSubmission)

	r.Get("/participants/{id}/submission_counts", HandleGetParticipantSubmissionCounts)

	r.Post("/review/{submissionId}", HandlePostReview)

	return r
}
