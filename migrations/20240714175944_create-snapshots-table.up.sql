CREATE TABLE IF NOT EXISTS snapshots (
    snap_id serial PRIMARY KEY,
    snap_date varchar(10) CHECK (snap_date ~ '^\d{2}/\d{2}/\d{4}$') NOT NULL, -- MM/DD/YYYY,
    total numeric(13, 2) DEFAULT 0,
    user_id integer NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,

    FOREIGN KEY (user_id) REFERENCES users(user_id)
);