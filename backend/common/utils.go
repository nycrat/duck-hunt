package common

import (
	"backend/types"
	"crypto"
	"crypto/pbkdf2"
	"database/sql"
	"encoding/base64"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateJwtToken(jwtString string, key []byte) (int, bool) {
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

	return 0, false
}

func GenerateJwtToken(id int, key []byte) string {
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

func DbFetchParticipants(db *sql.DB) []types.Participant {
	rows, err := db.Query("SELECT id, name, score FROM participants")

	if err != nil {
		log.Fatal(err)
	}

	participants := []types.Participant{}

	for rows.Next() {
		var id int
		var name string
		var score int
		err := rows.Scan(&id, &name, &score)
		if err != nil {
			log.Fatal(err)
		}

		name = strings.TrimSpace(name)
		participants = append(participants, types.Participant{Id: id, Name: name, Score: score})
	}

	return participants
}

func DbFetchActivityPreviews(db *sql.DB) []types.ActivityPreview {
	rows, err := db.Query("SELECT title, points FROM activities")

	if err != nil {
		log.Fatal(err)
	}

	activities := []types.ActivityPreview{}

	for rows.Next() {
		var title string
		var points int
		err := rows.Scan(&title, &points)
		if err != nil {
			log.Fatal(err)
		}

		title = strings.TrimSpace(title)
		activities = append(activities, types.ActivityPreview{Title: title, Points: points})
	}

	return activities
}

func DbFetchActivity(db *sql.DB, title string) (types.Activity, bool) {
	var points int
	var description string
	err := db.QueryRow(`SELECT points, description FROM activities WHERE title = $1`, title).Scan(&points, &description)

	if err != nil {
		return types.Activity{}, false
	}

	return types.Activity{
		Title:       title,
		Points:      points,
		Description: description,
	}, true
}

func DbSelectId(passcode string, pepper []byte, db *sql.DB) (int, bool) {
	hashedPasscode, err := pbkdf2.Key(crypto.SHA256.New, passcode, pepper, 4096, 64)

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

func DbFetchSubmissions(db *sql.DB, id int, title string) []types.Submission {
	rows, err := db.Query(`SELECT status, image FROM submissions WHERE participant_id = $1 AND activity_title = $2`, id, title)

	if err != nil {
		log.Fatal(err)
	}

	submissions := []types.Submission{}

	for rows.Next() {
		var status string
		var image []byte
		err := rows.Scan(&status, &image)
		if err != nil {
			log.Fatal(err)
		}

		submissions = append(submissions, types.Submission{Status: status, Image: image})
	}

	return submissions
}

func DbPostNewSubmission(db *sql.DB, id int, title string, image []byte) {
	_, err := db.Query(`INSERT INTO submissions (participant_id, activity_title, image) VALUES($1, $2, $3)`, id, title, image)

	if err != nil {
		log.Fatal(err)
	}
}

func DbFetchParticipantById(db *sql.DB, id int) types.Participant {
	var name string
	var score int

	err := db.QueryRow(`SELECT name, score FROM participants WHERE id = $1`, id).Scan(&name, &score)

	if err != nil {
		log.Fatal(err)
	}

	return types.Participant{Id: id, Name: name, Score: score}
}
