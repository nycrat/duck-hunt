package main

import (
	"backend/api"
	"backend/common"

	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Not enough arguments please specify JWT_HS256_KEY PEPPER DATABASE_URL")
	}

	hs256Key := []byte(os.Args[1])
	pepper := []byte(os.Args[2])
	dbConnStr := os.Args[3]
	db, err := sql.Open("postgres", dbConnStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*"},

		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(common.GetJwtMiddleware(hs256Key))
	r.Use(common.GetPepperMiddleware(pepper))
	r.Use(common.GetKeyMiddleware(hs256Key))
	r.Use(common.GetDbMiddleware(db))

	r.Post("/auth", api.HandlePostAuth)
	r.Post("/auth/admin", api.HandlePostAuthAdmin)
	r.Post("/session", api.HandlePostSession)

	r.Get("/participants", api.HandleGetParticipants)
	r.Get("/participants/{id}", api.HandleGetParticipantInfo)
	r.Get("/activities", api.HandleGetActivityPreviews)
	r.Get("/activities/{title}", api.HandleGetActivity)

	r.Get("/submissions/{title}", api.HandleGetSubmissions)
	r.Get("/submissions/{title}/{id}", api.HandleGetSubmissions)

	r.Post("/submissions/{title}", api.HandlePostSubmission)

	r.Get("/participants/{id}/submission_counts", api.HandleGetParticipantSubmissionCounts)

	r.Post("/review/{submissionId}", api.HandlePostReview)

	log.Println("Launched go web server on :8000")
	http.ListenAndServe(":8000", r)
}
