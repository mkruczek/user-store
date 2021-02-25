CREATE TABLE users
(
    id          UUID PRIMARY KEY,
    first_name  VARCHAR(50)         NOT NULL,
    last_name   VARCHAR(50)         NOT NULL,
    email       VARCHAR(255) UNIQUE NOT NULL,
    create_date BIGINT              NOT NULL
);