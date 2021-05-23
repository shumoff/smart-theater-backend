CREATE TABLE users
(
    id        serial PRIMARY KEY,
    username  VARCHAR(255) UNIQUE NOT NULL,
    password  VARCHAR(255)        NOT NULL,
    email     VARCHAR(511) UNIQUE NOT NULL,
    photo_url VARCHAR(255)        NULL,
    is_active boolean             NOT NULL
);
