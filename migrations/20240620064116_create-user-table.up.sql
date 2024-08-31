CREATE TYPE role_type AS ENUM (
    'DEFAULT',
    'ADMIN'
);

CREATE TABLE IF NOT EXISTS users (
  user_id serial PRIMARY KEY,
  email varchar(128) UNIQUE NOT NULL,
  enc_password varchar(256) NOT NULL,
  user_role role_type DEFAULT 'DEFAULT',
  last_logged_in timestamp DEFAULT '1970-01-01 00:00:00', -- Unix epoch
  created_at timestamp DEFAULT current_timestamp,
  updated_at timestamp DEFAULT current_timestamp
);

-- Create a case-insensitive unique index on the `email` column
CREATE UNIQUE INDEX unique_users_email ON users (LOWER(email));