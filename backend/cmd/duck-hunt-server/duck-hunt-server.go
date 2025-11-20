package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/nycrat/duck-hunt/backend/internal/api"
)

func main() {
	jwtKey := []byte(os.Getenv("DH_JWT_KEY"))
	pepper := []byte(os.Getenv("DH_PEPPER"))
	dbConnUrl := os.Getenv("DH_DATABASE_URL")

	if jwtKey == nil || pepper == nil || dbConnUrl == "" {
		log.Fatal("Environment variables not set: DH_JWT_KEY DH_PEPPER DH_DATABASE_URL")
	}

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
