package configs

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewConnectionDB() (*sql.DB, error) {
	coonStr := "password=root user=postgres dbname=postgres sslmode=disable"
	conn, err := sql.Open("postgres", coonStr)

	if err != nil {
		log.Fatal("Error connecting to the database")
	}

	return conn, err
}
