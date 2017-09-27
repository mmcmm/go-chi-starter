package main

import (
	"net/http"

	"github.com/mtdx/keyc/rest"
)

func main() {
	// TODO: uncomment
	// dbconn := db.Open()
	// defer dbconn.Close()

	r := rest.StartRouter()
	http.ListenAndServe(":3000", r)
}
