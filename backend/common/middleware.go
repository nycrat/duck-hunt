package common

import (
	"context"
	"database/sql"
	"net/http"
	"strings"
)

func GetJwtMiddleware(hs256Key []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			scheme, tokenString, found := strings.Cut(r.Header.Get("Authorization"), " ")

			if !found || scheme != "Bearer" {
				next.ServeHTTP(w, r)
				return
			}

			id, ok := ValidateJwtToken(tokenString, hs256Key)

			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), "id", id)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetDbMiddleware(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetPepperMiddleware(pepper []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "pepper", pepper)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetKeyMiddleware(hs256Key []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "key", hs256Key)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
