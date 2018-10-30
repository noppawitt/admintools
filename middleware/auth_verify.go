package middleware

import (
	"context"
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
			fmt.Println(authHeader)
			tokenString := strings.Split(authHeader, "Bearer")

			var err error
			if len(tokenString) > 1 {
				var token *jwt.Token
				token, err = jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					return secret, nil
				})

				if claims, ok := token.Claims.(model.TokenClaims); ok && token.Valid {
					ctx := context.WithValue(r.Context(), contextKey("id"), claims.ID)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
			}
			if err != nil {
				fmt.Println(err)
			}
			w.WriteHeader(http.StatusUnauthorized)
		})
	}
}
