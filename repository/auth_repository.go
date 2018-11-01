package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/noppawitt/admintools/config"
	"github.com/noppawitt/admintools/model"
)

// AuthRepository provides access an authentication external api (SSO)
type AuthRepository interface {
	// GetTokenByCode(code string, redirectURL string, consumerKey string, consumerSecret string) (*model.SSOToken, error)
	// GetTokenByRefreshToken(refreshToken string, consumerKey string, consumerSecret string) (*model.SSOToken, error)
	// GetUserInfo(accessToken string, consumerKey string, consumerSecret string) (*model.UserInfo, error)
	GetLogOut(ssoAccessToken string, redirectURL string) error
}

type authAgent struct {
	cfg *config.Config
}

// NewAuthAgent returns an auth agent
func NewAuthAgent(cfg *config.Config) AuthRepository {
	return &authAgent{cfg}
}

func (a *authAgent) GetLogOut(ssoAccessToken string, redirectURL string) error {
	req, _ := http.NewRequest("GET", a.cfg.AuthURL+"/auth/logout?key="+ssoAccessToken+"&redirectURL="+redirectURL, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case http.StatusSeeOther:
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		ssoToken := &model.SSOToken{}
		json.Unmarshal(respBody, ssoToken)
		return nil
	default:
		return errors.New("invalid access token")
	}
}
