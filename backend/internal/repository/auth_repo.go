package repository

import (
	"crypto"
	"crypto/pbkdf2"
	"database/sql"
	"encoding/base64"
	"log"
)

type AuthRepo struct {
	db     *sql.DB
	pepper []byte
}

type AuthRepoInterface interface {
	GetAuthorizedId(passcode string) (int, bool)
}

func NewAuthRepo(db *sql.DB, pepper []byte) *AuthRepo {
	return &AuthRepo{
		db:     db,
		pepper: pepper,
	}
}

func getEncodedHashedPasscode(passcode string, pepper []byte) (string, bool) {
	hashedPasscode, err := pbkdf2.Key(crypto.SHA256.New, passcode, pepper, 4096, 64)

	if err != nil {
		log.Println(err)
		return "", false
	}

	encodedHashedPasscode := base64.StdEncoding.EncodeToString(hashedPasscode)
	return encodedHashedPasscode, true
}

func (a *AuthRepo) GetAuthorizedId(passcode string) (int, bool) {
	encodedHashedPasscode, ok := getEncodedHashedPasscode(passcode, a.pepper)

	if !ok {
		return 0, false
	}

	var id int
	err := a.db.QueryRow(`SELECT participant_id FROM passcodes WHERE passcode = $1`, encodedHashedPasscode).Scan(&id)

	if err != nil {
		return 0, false
	}

	return id, true
}

func (a *AuthRepo) AddNewLoginInfo(id int, passcode string) {
	encodedHashedPasscode, ok := getEncodedHashedPasscode(passcode, a.pepper)

	if !ok {
		return
	}

	_, err := a.db.Query(`
	INSERT INTO passcodes (participant_id, passcode) VALUES ($1, $2)
	ON CONFLICT (participant_id) DO UPDATE
	SET passcode = $2
	`, id, encodedHashedPasscode)

	if err != nil {
		log.Println(err)
	}
}
