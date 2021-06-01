CREATE TABLE ratings
(
    user_id   integer REFERENCES users,
    movie_id  integer REFERENCES movies,
    rating    integer                             NOT NULL CHECK (rating > 0),
    timestamp timestamp default CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT pk_ratings PRIMARY KEY (user_id, movie_id)
);
