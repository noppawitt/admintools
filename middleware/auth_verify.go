package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/noppawitt/admintools/util"
)

type contextKey string

// AuthVerify verify access token from incomming requests
func AuthVerify(secret string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			tokenString := strings.Split(authHeader, "Bearer ")

			claims, err := util.ValidateToken(tokenString[1], secret)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), contextKey("User"), claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
