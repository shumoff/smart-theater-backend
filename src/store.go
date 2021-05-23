package main

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func (store *Store) GetMovieInfo(id string) (*Movie, error) {
	queryString := "SELECT id, title, description, release_date, poster_url from movies WHERE id = $1;"
	row := store.db.QueryRow(queryString, id)

	movie := &Movie{}

	err := row.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.PosterUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("there is no such movie: %w", err)
		}
		return nil, fmt.Errorf("could not get movie from db: %w", err)
	}

	return movie, nil
}
