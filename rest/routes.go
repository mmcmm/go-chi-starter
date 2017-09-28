package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/mtdx/keyc/config"
	"github.com/mtdx/keyc/rest/handler"
)

var r *chi.Mux
var tokenAuth *jwtauth.JwtAuth

func addRoutes() {
	r.Get("/", handler.Home)
	r.Get("/auth", handler.Auth)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/authenticated", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})
	})
}

// StartRouter create chi router & add the routes
func StartRouter() *chi.Mux {
	tokenAuth = jwtauth.New("HS256", []byte(config.JwtKey()), nil)
	r = chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	addRoutes()

	return r
}
