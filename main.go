package main

import (
	"net/http"

	"github.com/mtdx/keyc/db"
	"github.com/mtdx/keyc/rest"
)

func main() {
	dbconn := db.Open()
	defer dbconn.Close()

	r := rest.StartRouter(dbconn)
	http.ListenAndServe(":3000", r)
}
