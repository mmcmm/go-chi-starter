CREATE TABLE user
(
  id         BIGINT            NOT NULL
    CONSTRAINT user_pkey
    PRIMARY KEY,
  username   VARCHAR(255)
);