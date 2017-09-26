package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // postgress db driver
	"github.com/mtdx/keyc/config"
)

const (
	port = 5432
)

// Open a connection the db
func Open() *sql.DB {
	dbconfig := config.Env()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbconfig["host"], port, dbconfig["user"], dbconfig["password"], dbconfig["user"])

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open a db connection: %v\n", err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping the db: %v\n", err)
		os.Exit(1)
	}

	return db
}
