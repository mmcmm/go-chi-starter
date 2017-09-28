CREATE TABLE users
(
  id         BIGINT            NOT NULL
    CONSTRAINT user_pkey
    PRIMARY KEY,
  username   VARCHAR(255)
);