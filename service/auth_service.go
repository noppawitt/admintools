package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/noppawitt/admintools/util"

	"github.com/noppawitt/admintools/repository"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/noppawitt/admintools/config"
	"github.com/noppawitt/admintools/model"
)

// AuthService provides authentication service
type AuthService interface {
	AuthByCode(authRequest *model.AuthRequest) (*model.Token, error)
	AuthByRefreshToken(authRequest *model.AuthRequest) (*model.Token, error)
	LogOut(accessToken string, ssoAccessToken string, redirectURL string) error
	generateToken(id string, expiresIn time.Duration) (string, error)
	GenerateAccessToken(id string, expiresIn time.Duration) (string, error)
	GenerateRefreshToken(id string) (string, error)
	getUserInfo(accessToken string, consumerKey string) (*model.UserInfo, error)
}

type authService struct {
	authAgent repository.AuthRepository
	userRepo  repository.UserRepository
	cfg       *config.Config
}

// NewAuthService returns auth service
func NewAuthService(authAgent repository.AuthRepository, userRepo repository.UserRepository, cfg *config.Config) AuthService {
	return &authService{authAgent, userRepo, cfg}
}

func (s *authService) AuthByCode(authRequest *model.AuthRequest) (*model.Token, error) {
	req, _ := http.NewRequest("GET", s.cfg.AuthURL+"/auth/accesstoken", nil)
	req.Header.Add("code", authRequest.Code)
	req.Header.Add("redirectURL", authRequest.RedirectURL)
	req.Header.Add("consumerKey", authRequest.ConsumerKey)
	req.Header.Add("consumerSecret", s.cfg.ConsumerSecret)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		ssoToken := &model.SSOToken{}
		json.Unmarshal(respBody, ssoToken)

		// Create user if not exists and store token in database
		userInfo, err := s.getUserInfo(ssoToken.AccessToken, authRequest.ConsumerKey)
		if err != nil {
			return nil, err
		}

		user, err := s.userRepo.FindOne(userInfo.ID)
		if err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				user.ID = userInfo.ID
				user.FirstName = userInfo.FirstName
				user.LastName = userInfo.LastName
				user.SSORefreshToken = ssoToken.RefreshToken
				err = s.userRepo.Create(user)
				if err != nil {
					return nil, err
				}
			default:
				return nil, err
			}
		} else {
			user.SSORefreshToken = ssoToken.RefreshToken
		}

		// FIXME: move time to config
		accessToken, err := s.GenerateAccessToken(user.ID, 3600)
		if err != nil {
			panic(err)
		}
		refreshToken, err := s.GenerateRefreshToken(user.ID)
		if err != nil {
			panic(err)
		}

		user.RefreshToken = refreshToken
		err = s.userRepo.Save(user)
		if err != nil {
			return nil, err
		}

		token := &model.Token{
			AccessToken:    accessToken,
			RefreshToken:   refreshToken,
			Type:           "Bearer",
			SSOAccessToken: ssoToken.AccessToken,
		}

		return token, nil

	case http.StatusUnauthorized:
		return nil, errors.New("Invalid code")
	}

	return nil, errors.New("Cannot connect to SSO's server")
}

func (s *authService) AuthByRefreshToken(authRequest *model.AuthRequest) (*model.Token, error) {
	claims, err := util.ValidateToken(authRequest.RefreshToken, s.cfg.Secret)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.FindOne(claims.ID)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", s.cfg.AuthURL+"/auth/refresh", nil)
	req.Header.Add("refreshToken", user.SSORefreshToken)
	req.Header.Add("consumerKey", authRequest.ConsumerKey)
	req.Header.Add("consumerSecret", s.cfg.ConsumerSecret)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		ssoToken := &model.SSOToken{}
		json.Unmarshal(respBody, ssoToken)

		user.SSORefreshToken = ssoToken.RefreshToken
		err = s.userRepo.Save(user)
		if err != nil {
			return nil, err
		}

		// FIXME: move time to config
		accessToken, err := s.GenerateAccessToken(user.ID, 3600)
		if err != nil {
			return nil, err
		}
		token := &model.Token{
			AccessToken:    accessToken,
			Type:           "Bearer",
			SSOAccessToken: ssoToken.AccessToken,
		}

		return token, err
	case http.StatusUnauthorized:
		return nil, errors.New("Invalid code")
	}

	return nil, errors.New("Cannot connect to SSO's server")
}

func (s *authService) LogOut(accessToken string, ssoAccessToken string, redirectURL string) error {
	err := s.authAgent.GetLogOut(ssoAccessToken, redirectURL)
	if err != nil {
		return err
	}
	claims, err := util.ValidateToken(accessToken, s.cfg.Secret)
	if err != nil {
		return err
	}
	user, err := s.userRepo.FindOne(claims.ID)
	if err != nil {
		return err
	}
	user.RefreshToken = ""
	user.SSORefreshToken = ""
	s.userRepo.Save(user)
	return nil
}

func (s *authService) generateToken(id string, expiresIn time.Duration) (string, error) {
	expiresAt := int64(0)
	now := time.Now()
	if expiresIn > 0 {
		expiresAt = now.Add(time.Second * expiresIn).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.TokenClaims{
		ID:   id,
		Type: "Bearer",
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: expiresAt,
		},
	})
	return token.SignedString([]byte(s.cfg.Secret))
}

func (s *authService) GenerateAccessToken(id string, expiresIn time.Duration) (string, error) {
	return s.generateToken(id, expiresIn)
}

func (s *authService) GenerateRefreshToken(id string) (string, error) {
	return s.generateToken(id, 0)
}

func (s *authService) getUserInfo(accessToken string, consumerKey string) (*model.UserInfo, error) {
	req, _ := http.NewRequest("GET", s.cfg.AuthURL+"/user/info", nil)
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("consumerKey", consumerKey)
	req.Header.Add("consumerSecret", s.cfg.ConsumerSecret)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		userInfo := &model.UserInfo{}
		json.Unmarshal(respBody, &userInfo)
		return userInfo, nil
	default:
		return nil, err
	}
}
