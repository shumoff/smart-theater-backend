package main

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func (store *Store) GetMovieInfo(id int32) (*Movie, error) {
	queryString := "SELECT id, title, description, release_date, poster_url from movies WHERE id = $1;"
	row := store.db.QueryRow(queryString, id)

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

func (store *Store) GetMoviesPreviewInfo(ids []int32) ([]*MoviePreview, error) {
	queryString := "SELECT id, title, description, release_date, poster_url from movies WHERE id = any($1)"

	rows, err := store.db.Query(queryString, pq.Array(ids))
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
