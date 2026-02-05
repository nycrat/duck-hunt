package repository

import (
	"log"
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

func GenerateJwtToken(id int, key []byte) (string, bool) {
	duration := 7 * 24 * 60 * 60 * 1000 * 1000 * 1000 // 7 days in nanoseconds
	expirationTime := time.Now().Add(time.Duration(duration)).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": id,
			"exp": expirationTime,
		})

	signedToken, err := t.SignedString(key)

	if err != nil {
		log.Println(err)
		return "", false
	}

	return signedToken, true
}
