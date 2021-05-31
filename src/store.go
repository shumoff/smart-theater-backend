package main

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func (s *Store) GetUserProfile(id int32) (*UserProfile, error) {
	queryString := "SELECT id, username, photo_url from users WHERE id = $1;"
	row := s.db.QueryRow(queryString, id)

	userProfile := &UserProfile{}

	err := row.Scan(&userProfile.UserId, &userProfile.Username, &userProfile.PhotoUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("there is no such user: %w", err)
		}
		return nil, fmt.Errorf("could not get user from db: %w", err)
	}

	return userProfile, nil
}

func (s *Store) GetUserMovies(userId int32, pagination *Pagination) ([]*MoviePreview, error) {
	queryString := `
	select m.id, m.title, m.description, m.release_date, m.poster_url 
	from ratings r
	    join movies m on r.movie_id = m.id
	where r.user_id = $1
	offset $2 limit $3;
	`

	rows, err := s.db.Query(queryString, userId, pagination.Offset, pagination.Limit)
	if err != nil {
		return nil, fmt.Errorf("could not get user movies from db: %w", err)
	}
	defer rows.Close()

	var userMovies []*MoviePreview

	for rows.Next() {
		movie := &MoviePreview{}
		err := rows.Scan(&movie.MovieId, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.PosterUrl)
		if err != nil {
			return nil, fmt.Errorf("could not get movies from db: %w", err)
		}
		userMovies = append(userMovies, movie)
	}

	return userMovies, nil
}

func (s *Store) GetMovieInfo(id int32) (*Movie, error) {
	queryString := "SELECT id, title, description, release_date, poster_url from movies WHERE id = $1;"
	row := s.db.QueryRow(queryString, id)

	movie := &Movie{}

	err := row.Scan(&movie.MovieId, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.PosterUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("there is no such movie: %w", err)
		}
		return nil, fmt.Errorf("could not get movie from db: %w", err)
	}

	return movie, nil
}

func (s *Store) RateMovie(userId int32, movieId int32, rating int32) error {
	queryString := `
	insert into ratings (user_id, movie_id, rating)
	values ($1, $2, $3)
	on conflict (user_id, movie_id) do update set rating = EXCLUDED.rating, timestamp = now()
	`
	_, err := s.db.Exec(queryString, userId, movieId, rating)
	if err != nil {
		return fmt.Errorf("unable to rate movie: %w", err)
	}

	return nil
}

func (s *Store) GetMoviesPreview(pagination *Pagination) ([]*MoviePreview, error) {
	queryString := `
	SELECT id, title, description, release_date, poster_url
	from movies m
	    join ratings r on m.id = r.movie_id
	group by m.id
	order by avg(r.rating) desc
	offset $1 limit $2;
	`

	rows, err := s.db.Query(queryString, pagination.Offset, pagination.Limit)
	if err != nil {
		return nil, fmt.Errorf("could not get movies from db: %w", err)
	}
	defer rows.Close()

	var movies []*MoviePreview

	for rows.Next() {
		movie := &MoviePreview{}
		err := rows.Scan(&movie.MovieId, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.PosterUrl)
		if err != nil {
			return nil, fmt.Errorf("could not get movies from db: %w", err)
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (s *Store) GetCertainMoviesPreview(ids []int32) ([]*MoviePreview, error) {
	queryString := `
	SELECT id, title, description, release_date, poster_url
	from movies 
	    join unnest($1::int[]) with ordinality t(id, ord) using (id)
	order by t.ord;
	`
	rows, err := s.db.Query(queryString, pq.Array(ids))
	if err != nil {
		return nil, fmt.Errorf("could not get movies from db: %w", err)
	}
	defer rows.Close()

	var movies []*MoviePreview

	for rows.Next() {
		movie := &MoviePreview{}
		err := rows.Scan(&movie.MovieId, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.PosterUrl)
		if err != nil {
			return nil, fmt.Errorf("could not get movies from db: %w", err)
		}
		movies = append(movies, movie)
	}

	return movies, nil
}
