CREATE TYPE asset_type AS ENUM (
    'CASH',
    'BILLS',
    'STB',
    'ITB',
    'LTB',
    'COMMODITIES',
    'GOLD',
    'REITS',
    'TSM',
    'DLCB',
    'DLCG',
    'DLCV',
    'DMCB',
    'DMCG',
    'DMCV',
    'DSCG',
    'DSCB',
    'DSCV',
    'ILCB',
    'ILCG',
    'ILCV',
    'IMCB',
    'IMCG',
    'IMCV',
    'ISCB',
    'ISCG',
    'ISCV',
    'CRYPTO',
    'OTHER'
);

CREATE TABLE IF NOT EXISTS holdings (
    holding_id serial PRIMARY KEY,
    name varchar(256) NOT NULL,
    ticker varchar(32),
    asset_category asset_type NOT NULL,
    expense_ratio_pct numeric(3,2),
    maturation_date varchar(10) CHECK (maturation_date ~ '^(\d{2}/\d{2}/\d{4})?$'), -- MM/DD/YYYY
    interest_rate_pct numeric(3,2),
    is_deprecated boolean NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,

    FOREIGN KEY (user_id) REFERENCES users(user_id)
);