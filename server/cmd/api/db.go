package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connectionString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *application) connectToDB() (*sql.DB, error) {
	connection, err := openDB(app.ConnectionString)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to DB")
	return connection, nil
}
