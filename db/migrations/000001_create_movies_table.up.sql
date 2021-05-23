CREATE TABLE movies
(
    id          serial PRIMARY KEY,
    title       VARCHAR(255)   NOT NULL,
    description VARCHAR(16383) NULL,
    release_date integer        NULL,
    poster_url  VARCHAR(1023)  NULL,
    is_active   boolean        NOT NULL
);
