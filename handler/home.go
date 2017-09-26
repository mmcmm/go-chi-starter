package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type homeResponse struct {
	Elapsed int64 `json:"elapsed"`
}

func (rd *homeResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	rd.Elapsed = 10
	return nil
}

// Home route handler
func Home(w http.ResponseWriter, r *http.Request) {
	resp := &homeResponse{}

	render.Status(r, http.StatusOK)
	render.Render(w, r, resp)
}
