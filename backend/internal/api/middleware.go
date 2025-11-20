package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/nycrat/duck-hunt/backend/internal/repository"
)

func GetJwtMiddleware(hs256Key []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			scheme, tokenString, found := strings.Cut(r.Header.Get("Authorization"), " ")

			if !found || scheme != "Bearer" {
				next.ServeHTTP(w, r)
				return
			}

			id, ok := repository.ValidateJwtToken(tokenString, hs256Key)

			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), "id", id)
			ctx = context.WithValue(ctx, "admin", id == 1)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
