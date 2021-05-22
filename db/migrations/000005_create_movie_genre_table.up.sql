CREATE TABLE movie_genre
(
    genre_id integer REFERENCES genres,
    movie_id integer REFERENCES movies
);
