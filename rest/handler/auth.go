package handler

import (
	"net/http"
	"time"

	"github.com/mtdx/keyc/config"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

var tokenAuth *jwtauth.JwtAuth
var timeFunc func() time.Time

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
	origIat := time.Now().Unix()
	claims := jwtauth.Claims{"id": 123, "exp": exp, "orig_iat": origIat}
	_, tokenString, _ := tokenAuth.Encode(claims)
	resp := &authResponse{Token: tokenString}

	render.Status(r, http.StatusOK)
	render.Render(w, r, resp)
}
