package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/noppawitt/admintools/config"
	"github.com/noppawitt/admintools/model"
)

// AuthService provides authentication service
type AuthService interface {
	AuthByCode(authRequest *model.AuthRequest) (*model.Token, error)
	generateToken(id int, expiresIn time.Duration) (string, error)
	GenerateAccessToken(id int, expiresIn time.Duration) (string, error)
	GenerateRefreshToken(id int) (string, error)
}

type authService struct {
	cfg *config.Config
}

// NewAuthService returns auth service
func NewAuthService(cfg *config.Config) AuthService {
	return &authService{cfg}
}

func (s *authService) AuthByCode(authRequest *model.AuthRequest) (*model.Token, error) {
	req, _ := http.NewRequest("GET", s.cfg.AuthURL+"/auth/accesstoken", nil)
	req.Header.Add("code", authRequest.Code)
	req.Header.Add("consumerKey", authRequest.ConsumerKey)
	req.Header.Add("consumerSecret", s.cfg.ConsumerSecret)
	req.Header.Add("redirectURL", authRequest.RedirectURL)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		ssoToken := &model.SSOToken{}
		json.Unmarshal(respBody, &ssoToken)

		// TODO: implement this
		accessToken, err := s.GenerateAccessToken(1, 300)
		if err != nil {
			panic(err)
		}
		refreshToken, err := s.GenerateRefreshToken(1)
		if err != nil {
			panic(err)
		}
		token := &model.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			Type:         "Bearer",
			SSOToken:     ssoToken.AccessToken,
		}
		return token, err
	}

	return nil, errors.New("Cannot connect to SSO's server")
}

func (s *authService) generateToken(id int, expiresIn time.Duration) (string, error) {
	expiresAt := int64(0)
	now := time.Now()
	if expiresIn > 0 {
		expiresAt = now.Add(expiresIn).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, model.TokenClaims{
		ID:   id,
		Type: "Bearer",
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: expiresAt,
		},
	})
	return token.SignedString(s.cfg.Secret)
}

func (s *authService) GenerateAccessToken(id int, expiresIn time.Duration) (string, error) {
	return s.generateToken(id, expiresIn)
}

func (s *authService) GenerateRefreshToken(id int) (string, error) {
	return s.generateToken(id, 0)
}
