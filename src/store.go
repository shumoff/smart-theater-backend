package main

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func (store *Store) GetMovieInfo(id string) (*Movie, error) {
	rows, err := store.db.Query(
		"SELECT id, title, genre, description, release_date, poster_url from movies WHERE id = " +
			id + ";",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	movie := &Movie{}

	for rows.Next() {
		err = rows.Scan(movie.Id, movie.Title, movie.Genre, movie.Description, movie.ReleaseDate, movie.PosterUrl)
		if err != nil {
			return nil, fmt.Errorf("could not get movie from db: %w", err)
		}
	}

	return movie, nil
}
