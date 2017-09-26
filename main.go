package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mtdx/keyc/db"
)

func main() {
	dbconn := db.Open()
	defer dbconn.Close()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	http.ListenAndServe(":3000", r)
}
