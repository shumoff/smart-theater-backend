package main

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
	Rating int32 `json:"rating"`
}

type UserProfile struct {
	UserId   int32   `json:"user_id"`
	Username string  `json:"username"`
	PhotoUrl *string `json:"photo_url"`
}
