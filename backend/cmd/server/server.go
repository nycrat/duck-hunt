package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/nycrat/duck-hunt/backend/internal/api"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Not enough arguments please specify JWT_HS256_KEY PEPPER DATABASE_URL")
	}

	jwtKey := []byte(os.Args[1])
	pepper := []byte(os.Args[2])
	dbConnUrl := os.Args[3]

	// TODO: refactor this away from main function ?
	db, err := sql.Open("postgres", dbConnUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := api.DuckHuntRouter(jwtKey, pepper, db)

	log.Println("Launched go web server on :8000")
	http.ListenAndServe(":8000", router)
}
