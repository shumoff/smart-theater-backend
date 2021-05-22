package main

import (
	"database/sql"
)

type Store struct {
	db *sql.DB
}

//CREATE TABLE movies (
//	id SERIAL PRIMARY KEY,
//	title VARCHAR(256),
//	genre VARCHAR(256),
//	description VARCHAR(1024),
//	releaseDate integer,
//	poster_url VARCHAR(1024)
//);

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
			return nil, err
		}
	}

	return movie, nil
}
