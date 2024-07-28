CREATE TABLE IF NOT EXISTS snapshots (
    snap_id serial PRIMARY KEY,
    description text,
    snap_date varchar(10) CHECK (snap_date ~ '^\d{2}/\d{2}/\d{4}$') NOT NULL, -- MM/DD/YYYY,
    total numeric(13, 2) DEFAULT 0,
    weighted_er numeric(3, 3) DEFAULT 0,
    benchmark_id integer,
    user_id integer NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,

    FOREIGN KEY (benchmark_id) REFERENCES benchmarks(benchmark_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);