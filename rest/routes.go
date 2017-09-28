package rest

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/mtdx/keyc/rest/handler"
)

var r *chi.Mux

func addRoutes() {
	r.Get("/", handler.Home)
	r.Get("/auth", handler.Auth)
}

// StartRouter create chi router & add the routes
func StartRouter() *chi.Mux {
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
