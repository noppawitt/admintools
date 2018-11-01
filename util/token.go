package util

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/noppawitt/admintools/model"
)

// ValidateToken validates token and returns cliams
func ValidateToken(tokenString string, secret string) (*model.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*model.TokenClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
