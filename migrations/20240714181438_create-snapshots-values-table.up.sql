CREATE TABLE IF NOT EXISTS snapshots_values (
    snap_val_id serial PRIMARY KEY,
    snap_id integer NOT NULL,
    account_id integer NOT NULL,
    holding_id integer NOT NULL,
    total numeric(10, 2) NOT NULL,
    skip_rebalance boolean DEFAULT false,
    user_id integer NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,

    FOREIGN KEY (snap_id) REFERENCES snapshots(snap_id),
    FOREIGN KEY (account_id) REFERENCES accounts(account_id),
    FOREIGN KEY (holding_id) REFERENCES holdings(holding_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);