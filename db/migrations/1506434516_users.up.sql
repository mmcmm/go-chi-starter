CREATE TABLE users
(
  id         BIGINT            NOT NULL
    CONSTRAINT users_pkey
    PRIMARY KEY,
  username   VARCHAR(255)
);