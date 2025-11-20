package repository

import (
	"crypto"
	"crypto/pbkdf2"
	"database/sql"
	"encoding/base64"
	"log"
)

func DbSelectId(passcode string, pepper []byte, db *sql.DB) (int, bool) {
	hashedPasscode, err := pbkdf2.Key(crypto.SHA256.New, passcode, pepper, 4096, 64)

	if err != nil {
		log.Println(err)
		return 0, false
	}

	encodedHashedPasscode := base64.StdEncoding.EncodeToString(hashedPasscode)

	var id int
	err = db.QueryRow(`SELECT participant_id FROM passcodes WHERE passcode = $1`, encodedHashedPasscode).Scan(&id)

	if err != nil {
		return 0, false
	}

	return id, true
}
