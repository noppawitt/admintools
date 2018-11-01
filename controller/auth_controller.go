package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/noppawitt/admintools/model"
	"github.com/noppawitt/admintools/service"
)

// AuthController is an authentication controller
type AuthController struct {
	*Controller
	Service service.AuthService
}

const (
	codeGrantType         = "code"
	refreshTokenGrantType = "refreshToken"
)

// NewAuthController returns auth controller
func NewAuthController(s service.AuthService) *AuthController {
	return &AuthController{Service: s}
}

// Router is returns auth router
func (c *AuthController) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/token", c.token)
	r.Post("/logout", c.logOut)
	return r
}

func (c *AuthController) token(w http.ResponseWriter, r *http.Request) {
	var token *model.Token
	var err error

	request := &model.AuthRequest{}
	if err = json.NewDecoder(r.Body).Decode(request); err != nil {
		panic(err)
	}
	if request.GrantType == codeGrantType {
		token, err = c.Service.AuthByCode(request)
	} else if request.GrantType == refreshTokenGrantType {
		token, err = c.Service.AuthByRefreshToken(request)
	}
	if err != nil {
		panic(err)
	}
	c.JSON(w, http.StatusOK, token)
}

func (c *AuthController) logOut(w http.ResponseWriter, r *http.Request) {
	request := &model.Token{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		panic(err)
	}
	if err := c.Service.LogOut(request.AccessToken, request.SSOAccessToken, ""); err != nil {
		panic(err)
	}
	c.JSON(w, http.StatusOK, map[string]string{
		"message": "success"
	})
}
