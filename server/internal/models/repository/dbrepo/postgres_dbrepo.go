package dbrepo

import (
	"context"
	"database/sql"
	"server/internal/models"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 5 //5 secconds

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var query string = `
	select 
		id,
		title,
		release_date,
		runtime,
		mpaa_rating, 
		description,
		coalesce(image, ''),
		created_at,
		updated_at
		from movies
		order by title
	`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var movies []*models.Movie

	for rows.Next() {
		var movie models.Movie

		err := rows.Scan(
			&movie.Id,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.MPAARating,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}

func (m *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select id, email, first_name, last_name, password, created_at, updated_at from users where email = $1
	`
	var user models.User

	row := m.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (m *PostgresDBRepo) GetUserById(Id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select id, email, first_name, last_name, password, created_at, updated_at from users where id = $1
	`
	var user models.User

	row := m.DB.QueryRowContext(ctx, query, Id)

	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (m *PostgresDBRepo) GetMovieById(Id int) (*models.Movie, []*models.Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var query string = `
	select 
		id,
		title,
		release_date,
		runtime,
		mpaa_rating, 
		description,
		coalesce(image, ''),
		created_at,
		updated_at
		from movies
		where id = $1
	`

	var movie models.Movie

	row := m.DB.QueryRowContext(ctx, query, Id)

	err := row.Scan(
		&movie.Id,
		&movie.Title,
		&movie.ReleaseDate,
		&movie.Runtime,
		&movie.MPAARating,
		&movie.Description,
		&movie.Image,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		return nil, nil, err
	}


	query = `select g.id, g.genre from movies_genres mg
		left join genres g on mg.genre_id = g.id
		where mg.movie_id = $1
		order by g.genre
	`
	rows, err := m.DB.QueryContext(ctx, query, Id)

	if err != nil && err != sql.ErrNoRows {
		return nil, nil, err

	}
	defer rows.Close()

	var genres []*models.Genre
	var genresArray []int

	for rows.Next() {
		var genre models.Genre

		err := rows.Scan(
			&genre.Id,
			&genre.Genre,
		)

		if err != nil {
			return nil, nil, err
		}

		genres = append(genres, &genre)
		genresArray = append(genresArray, genre.Id)
	}


	movie.Genres = genres
	movie.GenresArray = genresArray
	allGenres, err := m.AllGenres()
	
	if err != nil {
		return nil, nil, err
	}

	return &movie, allGenres, nil

}

func (m *PostgresDBRepo) AllGenres() ([]*models.Genre, error){

	// id integer NOT NULL,
    // genre character varying(255),
    // created_at timestamp without time zone,
    // updated_at timestamp without time zone

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var query string = `
	select 
		id,
		genre,
		created_at,
		updated_at
	from genres
	order by genre
	`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var genres []*models.Genre

	for rows.Next() {
		var genre models.Genre

		err := rows.Scan(
			&genre.Id,
			&genre.Genre,
			&genre.CreatedAt,
			&genre.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		genres = append(genres, &genre)
	}

	return genres, nil
}

func (m *PostgresDBRepo) InsertMovie(movie models.Movie) (int, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var query string = `
	insert into movies (title, description, release_date, runtime, mpaa_rating, created_at, updated_at, image)
	values($1,$2,$3,$4,$5,$6,$7,$8) returning id
	`
	var newId int

	err:= m.DB.QueryRowContext(ctx, query, 
		movie.Title, 
		movie.Description, 
		movie.ReleaseDate, 
		movie.Runtime, 
		movie.MPAARating, 
		movie.CreatedAt, 
		movie.UpdatedAt, 
		movie.Image,
	).Scan(&newId)
	
	if err != nil {
		return 0, err
	}

	return newId, nil
}