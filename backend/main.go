package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

type participant struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}
	// jwtSecret := os.Args[1]
	dbConnStr := os.Args[2]
	db, err := sql.Open("postgres", dbConnStr)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, score FROM participants")

	if err != nil {
		log.Fatal(err)
	}

	participants := []participant{}

	for rows.Next() {
		var name string
		var score int
		err := rows.Scan(&name, &score)
		if err != nil {
			log.Fatal(err)
		}

		name = strings.TrimSpace(name)
		participants = append(participants, participant{Name: name, Score: score})
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
