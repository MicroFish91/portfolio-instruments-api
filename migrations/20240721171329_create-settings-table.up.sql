CREATE TABLE IF NOT EXISTS settings (
    settings_id serial PRIMARY KEY,
    reb_thresh_pct integer DEFAULT 10,
    vp_thresh_pct integer DEFAULT 10,
    vp_enabled boolean DEFAULT false, 
    user_id integer NOT NULL,
    benchmark_id integer,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,

    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (benchmark_id) REFERENCES benchmarks(benchmark_id) 
);