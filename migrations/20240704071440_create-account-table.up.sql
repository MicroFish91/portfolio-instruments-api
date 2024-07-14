CREATE TYPE shelter_type AS ENUM (
    'TAXABLE',
    'TRADITIONAL',
    'ROTH',
    'HSA',
    '529'
);

CREATE TABLE IF NOT EXISTS accounts (
    account_id serial PRIMARY KEY,
    name varchar(64) NOT NULL,
    description varchar(1024),
    tax_shelter shelter_type NOT NULL, 
    institution varchar(64) NOT NULL,
    is_deprecated boolean NOT NULL DEFAULT false,
    user_id integer NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,

    FOREIGN KEY (user_id) REFERENCES users(user_id)
);