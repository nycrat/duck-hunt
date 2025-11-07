package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type participant struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	participants := []participant{
		{Name: "Kel", Score: 1600},
		{Name: "Aly", Score: 1000},
		{Name: "Ado", Score: 1800},
		{Name: "Kiy", Score: 1200},
		{Name: "Bec", Score: 1300},
		{Name: "Kie", Score: 2600},
		{Name: "Luc", Score: 900},
	}
	serialized, _ := json.Marshal(participants)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*"},

		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(serialized)
	})

	fmt.Println("Launched go web server on :8000")
	http.ListenAndServe(":8000", r)
}
