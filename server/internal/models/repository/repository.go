package repository

import (
	"database/sql"
	"server/internal/models"
)

type DatabaseRepository interface {
	GetMovieById(Id int) (*models.Movie, []*models.Genre, error)
	AllMovies() ([]*models.Movie, error)
	InsertMovie(movie models.Movie) (int, error)
	Connection() *sql.DB
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(Id int) (*models.User, error)
	AllGenres() ([]*models.Genre, error)
}
