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

const (
	codeGrantType         = "code"
	refreshTokenGrantType = "refreshToken"
)

func (c *AuthController) token(w http.ResponseWriter, r *http.Request) {
	request := model.AuthRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		panic(err)
	}
	token, err := c.Service.AuthByCode(&request)
	if err != nil {
		panic(err)
	}
	c.JSON(w, http.StatusOK, token)
}
