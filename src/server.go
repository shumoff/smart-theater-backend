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

type Pagination struct {
	Offset int32 `schema:"offset"`
	Limit  int32 `schema:"limit"`
}

func makeHandler(fn func(http.ResponseWriter, *http.Request), allowedMethods map[string]bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		_, ok := allowedMethods[r.Method]
		if !ok {
			http.Error(w, fmt.Sprintf("%s method is not allowed", r.Method), http.StatusMethodNotAllowed)
			return
		}

		fn(w, r)
	}
}

func parsePagination(queryParams url.Values) (*Pagination, error) {
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

func (s *Server) getUserProfileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRegexp := regexp.MustCompile("^/api/v1/users/([0-9]+)/profile$")
		intUserId, err := strconv.Atoi(idRegexp.FindStringSubmatch(r.URL.Path)[1])
		if err != nil {
			http.Error(w, "Invalid user_id", http.StatusUnprocessableEntity)
			return
		}

		userId := int32(intUserId)

		userProfile, err := s.store.GetUserProfile(userId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(userProfile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) getUserMoviesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRegexp := regexp.MustCompile("^/api/v1/users/([0-9]+)/movies$")
		intUserId, err := strconv.Atoi(idRegexp.FindStringSubmatch(r.URL.Path)[1])
		if err != nil {
			http.Error(w, "Invalid user_id", http.StatusUnprocessableEntity)
			return
		}

		userId := int32(intUserId)

		pagination, err := parsePagination(r.URL.Query())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userMovies, err := s.store.GetUserMovies(userId, pagination)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(userMovies)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) getMoviesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagination, err := parsePagination(r.URL.Query())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		moviePreviews, err := s.store.GetMoviesPreview(pagination)
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
		idRegexp := regexp.MustCompile("^/api/v1/ratings/user/([0-9]+)/movie/([0-9]+)$")

		intUserId, err := strconv.Atoi(idRegexp.FindStringSubmatch(r.URL.Path)[1])
		if err != nil {
			http.Error(w, "Invalid user_id", http.StatusUnprocessableEntity)
			return
		}

		intMovieId, err := strconv.Atoi(idRegexp.FindStringSubmatch(r.URL.Path)[2])
		if err != nil {
			http.Error(w, "Invalid movie_id", http.StatusUnprocessableEntity)
			return
		}

		userId := int32(intUserId)
		movieId := int32(intMovieId)

		var rating MovieRating

		err = json.NewDecoder(r.Body).Decode(&rating)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.recommenderClient.SaveNewRating(userId, movieId, rating.Rating)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = s.store.RateMovie(userId, movieId, rating.Rating)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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

		pagination, err := parsePagination(r.URL.Query())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		relevantMoviesIds, err := s.recommenderClient.GetRelevantMovies(userId, pagination.Offset, pagination.Limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		moviePreviews, err := s.store.GetCertainMoviesPreview(relevantMoviesIds)
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

func (s *Server) getSimilarMoviesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRegexp := regexp.MustCompile("^/api/v1/recommendations/movie/([0-9]+)$")
		intMovieId, err := strconv.Atoi(idRegexp.FindStringSubmatch(r.URL.Path)[1])
		if err != nil {
			http.Error(w, "Invalid user_id", http.StatusUnprocessableEntity)
			return
		}

		movieId := int32(intMovieId)

		pagination, err := parsePagination(r.URL.Query())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		relevantMoviesIds, err := s.recommenderClient.GetSimilarMovies(movieId, pagination.Offset, pagination.Limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		moviePreviews, err := s.store.GetCertainMoviesPreview(relevantMoviesIds)
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

func (s *Server) getRecommenderSimilarMoviesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idRegexp := regexp.MustCompile("^/api/v1/recommendations/user/([0-9]+)/movie/([0-9]+)$")

		intUserId, err := strconv.Atoi(idRegexp.FindStringSubmatch(r.URL.Path)[1])
		if err != nil {
			http.Error(w, "Invalid user_id", http.StatusUnprocessableEntity)
			return
		}

		intMovieId, err := strconv.Atoi(idRegexp.FindStringSubmatch(r.URL.Path)[2])
		if err != nil {
			http.Error(w, "Invalid movie_id", http.StatusUnprocessableEntity)
			return
		}

		userId := int32(intUserId)
		movieId := int32(intMovieId)

		pagination, err := parsePagination(r.URL.Query())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		relevantMoviesIds, err := s.recommenderClient.GetRelevantSimilarMovies(userId, movieId, pagination.Offset, pagination.Limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		moviePreviews, err := s.store.GetCertainMoviesPreview(relevantMoviesIds)
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
	var rh RegexpHandler
	rh.HandleFunc(regexp.MustCompile("^/api/v1/users/[0-9]+/profile$"), makeHandler(s.getUserProfileHandler(), map[string]bool{"GET": true}))
	rh.HandleFunc(regexp.MustCompile("^/api/v1/users/[0-9]+/movies$"), makeHandler(s.getUserMoviesHandler(), map[string]bool{"GET": true}))

	rh.HandleFunc(regexp.MustCompile("^/api/v1/movies$"), makeHandler(s.getMoviesHandler(), map[string]bool{"GET": true}))
	rh.HandleFunc(regexp.MustCompile("^/api/v1/movies/[0-9]+$"), makeHandler(s.getMovieHandler(), map[string]bool{"GET": true}))

	// TODO SMART-12: implement elasticsearch
	//rh.HandleFunc(regexp.MustCompile("^/api/v1/movies/search$]"), makeHandler(s., map[string]bool{"GET": true}))

	rh.HandleFunc(regexp.MustCompile("^/api/v1/ratings/user/[0-9]+/movie/[0-9]+$"), makeHandler(s.rateMovieHandler(), map[string]bool{"POST": true}))

	rh.HandleFunc(regexp.MustCompile("^/api/v1/recommendations/user/[0-9]+$"), makeHandler(s.getRecommendedMoviesHandler(), map[string]bool{"GET": true}))
	rh.HandleFunc(regexp.MustCompile("^/api/v1/recommendations/movie/[0-9]+$"), makeHandler(s.getSimilarMoviesHandler(), map[string]bool{"GET": true}))
	rh.HandleFunc(regexp.MustCompile("^/api/v1/recommendations/user/[0-9]+/movie/[0-9]+$"), makeHandler(s.getRecommenderSimilarMoviesHandler(), map[string]bool{"GET": true}))

	http.HandleFunc("/", rh.ServeHTTP)

	log.Fatal(http.ListenAndServe(":"+s.port, nil))
}
