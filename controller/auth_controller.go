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
	} else {
		c.Error(w, http.StatusUnauthorized, "invalid grant type")
		return
	}
	if err != nil {
		c.Error(w, http.StatusUnauthorized, "failure")
		return
	}
	c.JSON(w, http.StatusOK, token)
}
