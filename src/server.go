package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
)

type Server struct {
	port  string
	store *Store
}

type Movie struct {
	Id          int64   `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	ReleaseDate *int16  `json:"release_date"`
	PosterUrl   *string `json:"poster_url"`
	IsActive    bool    `json:"is_active"`
}

func (s *Server) getMovieHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRegexp := regexp.MustCompile("^/(movies)/([0-9]+)$")
		movieId := idRegexp.FindStringSubmatch(r.URL.Path)[2]

		movie, err := s.store.GetMovieInfo(movieId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		err = json.NewEncoder(w).Encode(movie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) Serve() {
	http.HandleFunc("/movies/", s.getMovieHandler())

	log.Fatal(http.ListenAndServe(":"+s.port, nil))
}
