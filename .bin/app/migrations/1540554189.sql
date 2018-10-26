-- +migrate Up
CREATE TABLE cookie (
  id   BIGSERIAL NOT NULL,
  name TEXT      NOT NULL,
  CONSTRAINT pk_cookie_id PRIMARY KEY (id)
);

INSERT INTO cookie (name) VALUES ('Chocolate'), ('Shortbread');

-- +migrate Down
DROP TABLE cookie;
