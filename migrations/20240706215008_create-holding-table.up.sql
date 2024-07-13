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
    'ISCV'
);

CREATE TABLE IF NOT EXISTS holdings (
    holding_id serial PRIMARY KEY,
    name varchar(256) NOT NULL,
    ticker varchar(32),
    asset_category asset_type NOT NULL,
    expense_ratio real,
    maturation_date varchar(10) CHECK (maturation_date ~ '^(\d{2}/\d{2}/\d{4})?$'), -- MM/DD/YYYY
    interest_rate real,
    is_deprecated boolean NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp
);