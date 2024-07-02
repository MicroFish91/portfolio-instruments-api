CREATE TYPE tax_shelter AS ENUM (
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
    shelter_type tax_shelter NOT NULL, 
    institution varchar(64) NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,
    deleted_at timestamp
);