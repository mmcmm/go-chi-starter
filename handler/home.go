package handler

import "net/http"

// Home route handler
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}
