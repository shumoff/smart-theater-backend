CREATE TABLE ratings
(
    user_id   integer REFERENCES users,
    movie_id  integer REFERENCES movies,
    rating    integer   NOT NULL,
    timestamp timestamp NOT NULL
);
