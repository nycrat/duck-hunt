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

func (a *AuthRepo) GetAuthorizedId(passcode string) (int, bool) {
	hashedPasscode, err := pbkdf2.Key(crypto.SHA256.New, passcode, a.pepper, 4096, 64)

	if err != nil {
		log.Println(err)
		return 0, false
	}

	encodedHashedPasscode := base64.StdEncoding.EncodeToString(hashedPasscode)

	var id int
	err = a.db.QueryRow(`SELECT participant_id FROM passcodes WHERE passcode = $1`, encodedHashedPasscode).Scan(&id)

	if err != nil {
		return 0, false
	}

	return id, true
}
