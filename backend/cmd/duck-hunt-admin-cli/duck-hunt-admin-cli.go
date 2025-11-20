package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/nycrat/duck-hunt/backend/internal/repository"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	jwtKey := []byte(os.Getenv("DH_JWT_KEY"))
	pepper := []byte(os.Getenv("DH_PEPPER"))
	dbConnUrl := os.Getenv("DH_DATABASE_URL")

	if jwtKey == nil || pepper == nil || dbConnUrl == "" {
		log.Fatal("Environment variables not set: DH_JWT_KEY DH_PEPPER DH_DATABASE_URL")
	}

	id, err := strconv.ParseInt(os.Args[1], 10, 32)
	passcode := os.Args[2]

	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", dbConnUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	authRepository := repository.NewAuthRepo(db, pepper)
	authRepository.AddNewLoginInfo(int(id), passcode)

	log.Printf("INSERTED NEW PASSWORD INTO DATABASE FOR USER: %d", id)
}
