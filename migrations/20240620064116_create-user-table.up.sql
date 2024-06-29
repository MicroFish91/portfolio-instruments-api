CREATE TABLE IF NOT EXISTS users (
  user_id serial PRIMARY KEY,
  email varchar(128) UNIQUE NOT NULL,
  enc_password varchar(256) NOT NULL,
  created_at timestamp DEFAULT current_timestamp,
  updated_at timestamp DEFAULT current_timestamp
);