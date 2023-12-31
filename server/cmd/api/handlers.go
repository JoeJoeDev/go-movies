package main

import (
	"errors"
	"net/http"
	"server/internal/models"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"messge"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Pong",
		Version: "1.0.0",
	}
	_ = app.writeJSON(w, http.StatusOK, payload)

}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, movies)

}

func (app *application) AllGenres(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllGenres()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, movies)

}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// read payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	//validate user against db
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("Invald credentials"), http.StatusBadRequest)
		return
	}

	//check password
	valid, err := user.PasswordMatches(requestPayload.Password)

	if err != nil || !valid {
		app.errorJSON(w, errors.New("Invald credentials"), http.StatusBadRequest)
		return
	}

	//create jwt user
	u := jwtUser{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	tokens, err := app.auth.GenerateTokenPair(&u)

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	app.writeJSON(w, http.StatusAccepted, tokens)

}

func (app * application) refreshToken (w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		if cookie.Name == app.auth.CookieName {
			claims := &Claims{}
			refreshToken := cookie.Value

			//pase the token to get the value
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error){
				return []byte(app.JWTSecret), nil
			})

			if err != nil{
				app.errorJSON(w, errors.New("unauthorised"), http.StatusUnauthorized)
			}
		

			// get user id for the token claims
			userId, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJSON(w, errors.New("unauthorised"), http.StatusUnauthorized)
				return
			}

			user, err := app.DB.GetUserById(userId)

			if err != nil {
				app.errorJSON(w, errors.New("unauthorised"), http.StatusUnauthorized)
				return
			}

			u := jwtUser{
				Id: user.Id,
				FirstName: user.FirstName,
				LastName: user.LastName,
			}

			tokenPairs, err := app.auth.GenerateTokenPair(&u)

			if err != nil {
				app.errorJSON(w, errors.New("erro generating tokens"), http.StatusUnauthorized)
				return
			}
			
			http.SetCookie(w, app.auth.GetRefreshCookie(tokenPairs.RefreshToken))
			app.writeJSON(w, http.StatusOK, tokenPairs)
		}
	}
	
}

func (app * application) logout (w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, app.auth.GetExpiredRefreshCookie())
	w.WriteHeader(http.StatusAccepted)

}


func (app * application) MovieCatalog (w http.ResponseWriter, r *http.Request) {
	
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, movies)

}

func (app *application) GetMovieById(w http.ResponseWriter, r *http.Request) {
	id, err :=  strconv.Atoi(chi.URLParam(r, "id"))
	
	if err != nil {
		app.errorJSON(w, errors.New("invalid id"))
		return
	}

	
	movie, genres, err := app.DB.GetMovieById(id)

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var responseBody = struct {
		Movie *models.Movie `json:"movie"`
		Genres []*models.Genre `json:"genres"`
	}{
		Movie: movie,
		Genres: genres,
	}

	_ = app.writeJSON(w, http.StatusOK, responseBody)

}

func (app *application) InsertMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	err := app.readJSON(w,r,&movie)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	//try to get an image

	// handle genres

	resp := JSONResponse{ Error: false, Message: "movie updated"}

	app.writeJSON(w, http.StatusAccepted, resp)

}

