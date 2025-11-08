package main

import (
	"crypto"
	"crypto/pbkdf2"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
)

type participant struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type activity struct {
	Title  string `json:"title"`
	Points int    `json:"points"`
	Link   string `json:"link"`
}

func dbFetchParticipants(db *sql.DB) []participant {
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

	return participants
}

func dbFetchActivities(db *sql.DB) []activity {
	rows, err := db.Query("SELECT title, points, link FROM activities")

	if err != nil {
		log.Fatal(err)
	}

	activities := []activity{}

	for rows.Next() {
		var title string
		var points int
		var link string
		err := rows.Scan(&title, &points, &link)
		if err != nil {
			log.Fatal(err)
		}

		title = strings.TrimSpace(title)
		link = strings.TrimSpace(link)
		activities = append(activities, activity{Title: title, Points: points, Link: link})
	}

	return activities
}

func validateJwtToken(jwtString string, key []byte) (int, bool) {
	token, err := jwt.Parse(jwtString, func(t *jwt.Token) (any, error) {
		return key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		log.Println(err)
		return 0, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return int(claims["sub"].(float64)), true
	}

	// TODO implement checking exp

	return 0, false
}

func generateJwtToken(id int, key []byte) string {
	duration := 48 * 60 * 60 * 1000 * 1000 * 1000 // 2 days in nanoseconds
	expirationTime := time.Now().Add(time.Duration(duration)).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": id,
			"exp": expirationTime,
		})

	signedToken, err := t.SignedString(key)

	if err != nil {
		log.Fatal(err)
	}

	return signedToken
}

func dbSelectId(passcode string, salt []byte, db *sql.DB) (int, bool) {
	hashedPasscode, err := pbkdf2.Key(crypto.SHA256.New, passcode, salt, 4096, 64)

	if err != nil {
		log.Fatal(err)
	}

	encodedHashedPasscode := base64.StdEncoding.EncodeToString(hashedPasscode)

	var id int
	err = db.QueryRow(`SELECT participant_id FROM passcodes WHERE passcode = $1`, encodedHashedPasscode).Scan(&id)

	if err != nil {
		return 0, false
	}

	return id, true
}

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Not enough arguments please specify RSA_PRIVATE_KEY CONSTANT_SALT DATABASE_URL")
	}

	hs256Key := []byte(os.Args[1])
	constantSalt := []byte(os.Args[2])
	dbConnStr := os.Args[3]
	db, err := sql.Open("postgres", dbConnStr)

	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*"},

		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Post("/auth", func(w http.ResponseWriter, r *http.Request) {
		scheme, passcode, found := strings.Cut(r.Header.Get("Authorization"), " ")

		if !found || scheme != "Basic" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		id, ok := dbSelectId(passcode, constantSalt, db)

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := generateJwtToken(id, hs256Key)
		w.Write([]byte(token))
	})

	r.Get("/participants", func(w http.ResponseWriter, r *http.Request) {
		scheme, tokenString, found := strings.Cut(r.Header.Get("Authorization"), " ")

		if !found || scheme != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		_, ok := validateJwtToken(tokenString, hs256Key)

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		participants := dbFetchParticipants(db)
		serialized, _ := json.Marshal(participants)
		w.Write(serialized)
	})

	r.Get("/activities", func(w http.ResponseWriter, r *http.Request) {
		scheme, tokenString, found := strings.Cut(r.Header.Get("Authorization"), " ")

		if !found || scheme != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		_, ok := validateJwtToken(tokenString, hs256Key)

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		activities := dbFetchActivities(db)
		serialized, _ := json.Marshal(activities)
		w.Write(serialized)
	})

	log.Println("Launched go web server on :8000")
	http.ListenAndServe(":8000", r)
}
