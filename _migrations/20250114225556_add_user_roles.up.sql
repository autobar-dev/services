CREATE TABLE user_roles (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);

INSERT INTO user_roles (id, name) VALUES (1, 'user');
INSERT INTO user_roles (id, name) VALUES (2, 'admin');