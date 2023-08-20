CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE unfinished_registrations(
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),

  email TEXT UNIQUE NOT NULL,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  date_of_birth DATE NOT NULL,
  locale VARCHAR(9) NOT NULL,

  -- necessary for email confirmation
  email_confirmed BOOLEAN NOT NULL DEFAULT false,
  email_confirmation_code VARCHAR(64) UNIQUE NOT NULL,
  email_confirmation_code_issued_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,

  -- -- further steps must allow for null values
  -- -- step: "phone-number"
  -- phone_number_country_code VARCHAR(9) DEFAULT NULL,
  -- phone_number VARCHAR(32) DEFAULT NULL,
  --
  -- -- necessary for phone number confirmation
  -- phone_number_confirmed BOOLEAN DEFAULT false,
  -- phone_number_code VARCHAR(6) DEFAULT NULL,

  -- always there
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);
