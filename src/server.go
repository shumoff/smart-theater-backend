package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/gorilla/schema"

	"github.com/shumoff/smart-theater-backend/recommender_grpc"
)

var decoder = schema.NewDecoder()

type Server struct {
	port              string
	store             *Store
	recommenderClient *recommender_grpc.RecommenderClient
}

type Movie struct {
	MovieId     int64   `json:"movie_id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	ReleaseDate *int16  `json:"release_date"`
	PosterUrl   *string `json:"poster_url"`
	IsActive    bool    `json:"is_active,omitempty"`
}

type MoviePreview struct {
	MovieId     int64   `json:"movie_id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	ReleaseDate *int16  `json:"release_date"`
	PosterUrl   *string `json:"poster_url"`
	IsActive    bool    `json:"is_active,omitempty"`
}

type MovieRating struct {
	Rating int `json:"rating"`
}

type Pagination struct {
	Offset int32 `schema:"offset"`
	Limit  int32 `schema:"limit"`
}

func makeHandler(fn func(http.ResponseWriter, *http.Request), allowedMethods map[string]bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, ok := allowedMethods[r.Method]
		if !ok {
			http.Error(w, fmt.Sprintf("%s method is not allowed", r.Method), http.StatusMethodNotAllowed)
			return
		}

		fn(w, r)
	}
}

func parseOffsetLimit(queryParams url.Values) (*Pagination, error) {
	var pagination Pagination

	err := decoder.Decode(&pagination, queryParams)
	if err != nil {
		return nil, fmt.Errorf("invalid offset or limit: %w", err)
	}
	if pagination.Limit == 0 {
		pagination.Limit = 20
	}

	return &pagination, nil
}

func (s *Server) getMovieHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRegexp := regexp.MustCompile("^/api/v1/movies/([0-9]+)$")
		intMovieId, err := strconv.Atoi(idRegexp.FindStringSubmatch(r.URL.Path)[1])
		if err != nil {
			http.Error(w, "Invalid movie_id", http.StatusUnprocessableEntity)
			return
		}

		movieId := int32(intMovieId)

		movie, err := s.store.GetMovieInfo(movieId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(movie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) rateMovieHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rating MovieRating

		err := json.NewDecoder(r.Body).Decode(&rating)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (s *Server) getRecommendedMoviesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRegexp := regexp.MustCompile("^/api/v1/recommendations/user/([0-9]+)$")
		intUserId, err := strconv.Atoi(idRegexp.FindStringSubmatch(r.URL.Path)[1])
		if err != nil {
			http.Error(w, "Invalid user_id", http.StatusUnprocessableEntity)
			return
		}

		userId := int32(intUserId)

		pagination, err := parseOffsetLimit(r.URL.Query())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		relevantMoviesIds := s.recommenderClient.GetRelevantMovies(userId, pagination.Offset, pagination.Limit)

		moviePreviews, err := s.store.GetMoviesPreviewInfo(relevantMoviesIds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(moviePreviews)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) Serve() {
	http.HandleFunc("/api/v1/movies/", makeHandler(s.getMovieHandler(), map[string]bool{"GET": true}))
	http.HandleFunc("/api/v1/recommendations/user/", makeHandler(s.getRecommendedMoviesHandler(), map[string]bool{"GET": true}))

	log.Fatal(http.ListenAndServe(":"+s.port, nil))
}
