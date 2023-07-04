package main

import (
	"fmt"
	"log"
	"net/http"
	"server/internal/models/repository"
	"server/internal/models/repository/dbrepo"
	"time"
)

const port = 8080

type application struct {
	Domain           string
	ConnectionString string
	DB               repository.DatabaseRepository
	auth             Auth
	JWTSecret        string
	JWTIssuer        string
	JWTAudience      string
	CookieDomain     string
}

func main() {
	// set app config
	var app application
	// read from command line

	//conect to db
	app.ConnectionString = "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5"
	app.JWTSecret = "secret"
	app.JWTIssuer = "example.com"
	app.JWTAudience = "example.com"
	app.CookieDomain = "localhost"
	app.Domain = "example.com"

	conn, err := app.connectToDB()

	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()
	app.auth = Auth{
		Issuer: app.JWTIssuer,
		Audience: app.JWTAudience,
		Secret: app.JWTSecret,
		TokeExp: time.Minute * 16,
		RefreshExp: time.Hour * 24,
		CookiePath: "/",
		CookieName: "refresh_token",
		CookieDomain: app.CookieDomain,
	}
	log.Println("Running on port", port)

	//start webserver
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}

}
