package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/noppawitt/admintools/model"

	jwt "github.com/dgrijalva/jwt-go"
)

type contextKey string

// AuthVerify verify access token from incomming requests
func AuthVerify(secret string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			tokenString := strings.Split(authHeader, "Bearer ")

			token, err := jwt.ParseWithClaims(tokenString[1], &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(secret), nil
			})

			if claims, ok := token.Claims.(*model.TokenClaims); ok && token.Valid {
				fmt.Println(claims.ID)
				next.ServeHTTP(w, r)
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
			}
		})
	}
}
