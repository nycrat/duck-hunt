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

func DbFetchParticipants(db *sql.DB) []types.Participant {
	rows, err := db.Query("SELECT name, score FROM participants")

	if err != nil {
		log.Fatal(err)
	}

	participants := []types.Participant{}

	for rows.Next() {
		var name string
		var score int
		err := rows.Scan(&name, &score)
		if err != nil {
			log.Fatal(err)
		}

		name = strings.TrimSpace(name)
		participants = append(participants, types.Participant{Name: name, Score: score})
	}

	return participants
}

func DbFetchActivities(db *sql.DB) []types.Activity {
	rows, err := db.Query("SELECT title, points, link FROM activities")

	if err != nil {
		log.Fatal(err)
	}

	activities := []types.Activity{}

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
		activities = append(activities, types.Activity{Title: title, Points: points, Link: link})
	}

	return activities
}

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

	// TODO implement checking exp

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
