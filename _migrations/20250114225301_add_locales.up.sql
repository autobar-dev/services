CREATE TABLE locales (
  id SERIAL PRIMARY KEY,

  code TEXT UNIQUE NOT NULL,

  name TEXT NOT NULL,
  name_localized TEXT NOT NULL,

  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);

INSERT INTO locales (id, code, name, name_localized) VALUES (1, 'en', 'English', 'English');
INSERT INTO locales (id, code, name, name_localized) VALUES (2, 'pl', 'Polish', 'Polski');