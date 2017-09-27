package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type homeResponse struct {
	Elapsed int64 `json:"elapsed"`
}

func (rd *homeResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Home route handler
func Home(w http.ResponseWriter, r *http.Request) {
	resp := &homeResponse{}

	render.Status(r, http.StatusOK)
	render.Render(w, r, resp)
}
