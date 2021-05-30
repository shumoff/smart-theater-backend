CREATE TABLE movie_genre
(
    genre_id integer REFERENCES genres,
    movie_id integer REFERENCES movies,
    CONSTRAINT pk_movie_genre PRIMARY KEY (genre_id, movie_id)
);
