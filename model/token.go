package model

import (
	"github.com/dgrijalva/jwt-go"
)

const (
	// CodeGrantType is a code grant type
	CodeGrantType = "code"
	// RefreshTokenGrantType is a refresh token grant type
	RefreshTokenGrantType = "refreshToken"
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
	AccessToken    string `json:"accessToken"`
	RefreshToken   string `json:"refreshToken,omitempty"`
	Type           string `json:"type"`
	SSOAccessToken string `json:"ssoAccessToken"`
}

// TokenClaims is a token claims model
type TokenClaims struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	jwt.StandardClaims
}

// UserInfo is a SSO's user information model
type UserInfo struct {
	ID        string `json:"userID"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
