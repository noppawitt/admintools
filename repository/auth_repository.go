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
	GetSSOTokenByCode(code string, redirectURL string, consumerKey string) (*model.SSOToken, error)
	GetSSOTokenByRefreshToken(ssoRefreshToken string, consumerKey string) (*model.SSOToken, error)
	GetUserInfo(ssoAccessToken string, consumerKey string) (*model.UserInfo, error)
}

type authAgent struct {
	cfg *config.Config
}

// NewAuthAgent returns an auth agent
func NewAuthAgent(cfg *config.Config) AuthRepository {
	return &authAgent{cfg}
}

func (a *authAgent) GetSSOTokenByCode(code string, redirectURL string, consumerKey string) (*model.SSOToken, error) {
	req, _ := http.NewRequest("GET", a.cfg.AuthURL+"/auth/accesstoken", nil)
	req.Header.Add("code", code)
	req.Header.Add("redirectURL", redirectURL)
	req.Header.Add("consumerKey", consumerKey)
	req.Header.Add("consumerSecret", a.cfg.ConsumerSecret)
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

		return ssoToken, nil

	case http.StatusUnauthorized:
		return nil, errors.New("Invalid code")
	}

	return nil, errors.New("Cannot connect to SSO's server")
}

func (a *authAgent) GetSSOTokenByRefreshToken(ssoRefreshToken string, consumerKey string) (*model.SSOToken, error) {
	req, _ := http.NewRequest("GET", a.cfg.AuthURL+"/auth/refresh", nil)
	req.Header.Add("refreshToken", ssoRefreshToken)
	req.Header.Add("consumerKey", consumerKey)
	req.Header.Add("consumerSecret", a.cfg.ConsumerSecret)
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

		return ssoToken, nil
	case http.StatusUnauthorized:
		return nil, errors.New("Invalid code")
	}

	return nil, errors.New("Cannot connect to SSO's server")
}

func (a *authAgent) GetUserInfo(ssoAccessToken string, consumerKey string) (*model.UserInfo, error) {
	req, _ := http.NewRequest("GET", a.cfg.AuthURL+"/user/info", nil)
	req.Header.Add("Authorization", "Bearer "+ssoAccessToken)
	req.Header.Add("consumerKey", consumerKey)
	req.Header.Add("consumerSecret", a.cfg.ConsumerSecret)
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
