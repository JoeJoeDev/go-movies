package main

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Post("/authenticate", app.authenticate)
	mux.Get("/refresh", app.refreshToken)
	mux.Get("/logout", app.logout)

	mux.Get("/", app.Home)
	mux.Get("/movies", app.AllMovies)
	mux.Get("/movie/{id}", app.GetMovieById)

	mux.Get("/genres", app.AllGenres)

	mux.Route("/admin", func(mux chi.Router){
		mux.Use(app.authRequired)
		
		mux.Get("/movies", app.MovieCatalog)

		mux.Post("/movie", app.InsertMovie)
		
	})
	return mux
}
