CREATE TABLE movies
(
    id           serial PRIMARY KEY,
    title        VARCHAR(255)   NOT NULL,
    description  VARCHAR(16383) NULL,
    runtime      integer        NULL,
    release_date integer        NULL,
    poster_url   VARCHAR(1023)  NULL,
    is_active    boolean DEFAULT false
);
