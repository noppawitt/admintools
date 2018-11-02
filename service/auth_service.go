package service

import (
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
	LogOut(userID string) error
	generateToken(user *model.User, expiresIn int) (string, error)
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
	ssoToken, err := s.authAgent.GetSSOTokenByCode(authRequest.Code, authRequest.RedirectURL, authRequest.ConsumerKey)
	if err != nil {
		return nil, err
	}

	userInfo, err := s.authAgent.GetUserInfo(ssoToken.AccessToken, authRequest.ConsumerKey)
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
	}
	user.SSORefreshToken = ssoToken.RefreshToken

	accessToken, err := s.generateToken(user, s.cfg.AccessTokenExpiryTime)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.generateToken(user, s.cfg.RefreshTokenExpiryTime)
	if err != nil {
		return nil, err
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

	ssoToken, err := s.authAgent.GetSSOTokenByRefreshToken(user.SSORefreshToken, authRequest.ConsumerKey)
	if err != nil {
		// Force logout if can not refresh token from SSO
		s.LogOut(claims.ID)
		return nil, err
	}

	user.SSORefreshToken = ssoToken.RefreshToken
	err = s.userRepo.Save(user)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.generateToken(user, s.cfg.AccessTokenExpiryTime)
	if err != nil {
		return nil, err
	}
	token := &model.Token{
		AccessToken:    accessToken,
		Type:           "Bearer",
		SSOAccessToken: ssoToken.AccessToken,
	}

	return token, nil
}

func (s *authService) LogOut(userID string) error {
	user, err := s.userRepo.FindOne(userID)
	if err != nil {
		return err
	}
	user.RefreshToken = ""
	user.SSORefreshToken = ""
	s.userRepo.Save(user)
	return nil
}

func (s *authService) generateToken(user *model.User, expiresIn int) (string, error) {
	expiresAt := int64(0)
	now := time.Now()
	if expiresIn > 0 {
		expiresAt = now.Add(time.Second * time.Duration(expiresIn)).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.TokenClaims{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Type:      "Bearer",
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: expiresAt,
		},
	})
	return token.SignedString([]byte(s.cfg.Secret))
}
