package model

import (
	"github.com/dgrijalva/jwt-go"
)

// AuthRequest is an auth request model
type AuthRequest struct {
	Code         string `json:"code"`
	ConsumerKey  string `json:"consumerKey"`
	RedirectURL  string `json:"redirectURL"`
	RefreshToken string `json:"refreshToken"`
	GrantType    string `json:"grantType"`
}

// SSOToken is a sso's token model
type SSOToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenTypes   string `json:"tokenTypes"`
	ExpiresIn    int    `json:"expiresIn"`
	Scope        string `json:"scope"`
}

// Token is a token response model
type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Type         string `json:"type"`
	SSOToken     string `json:"ssoToken"`
}

// TokenClaims is a token claims model
type TokenClaims struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	jwt.StandardClaims
}
