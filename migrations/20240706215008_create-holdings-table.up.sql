CREATE TYPE asset_type AS ENUM (
    'TSM',
    'DLCB'
);

CREATE TABLE IF NOT EXISTS holdings (
    holding_id serial PRIMARY KEY,
    name varchar(256) NOT NULL,
    ticker varchar(32) NOT NULL,
    asset_category asset_type NOT NULL,
    expense_ratio real NOT NULL,
    is_deprecated boolean NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp
);