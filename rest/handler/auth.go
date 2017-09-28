package handler

import (
	"net/http"
	"time"

	"github.com/mtdx/keyc/config"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

var tokenAuth *jwtauth.JwtAuth

type authResponse struct {
	Token string `json:"token"`
}

func (rd *authResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Auth route handler
func Auth(w http.ResponseWriter, r *http.Request) {
	tokenAuth = jwtauth.New("HS256", []byte(config.JwtKey()), nil)
	exp := time.Now().Add(time.Hour * time.Duration(12)).Unix()
	claims := jwtauth.Claims{"user_id": 123, "exp": exp}
	_, tokenString, _ := tokenAuth.Encode(claims)
	resp := &authResponse{Token: tokenString}

	render.Status(r, http.StatusOK)
	render.Render(w, r, resp)
}
