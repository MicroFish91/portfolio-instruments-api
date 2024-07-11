CREATE TABLE IF NOT EXISTS benchmarks (
    benchmark_id serial PRIMARY KEY,
    name varchar(64) NOT NULL,
    description text,
    asset_allocation jsonb NOT NULL,
    std_dev_pct real,
    real_return_pct real,
    drawdown_yrs int,
    is_deprecated boolean NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp
);