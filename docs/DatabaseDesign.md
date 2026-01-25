## 🗄 Database Design

- Using _'s in model props so as to correspond directly with the SQL database values.  This allows us to easily marshal values using `row.Scan()` in our store logic.

#### Users: 

    | Column Name    | Datatype  | Not Null | Primary Key |
    |----------------|-----------|----------|-------------|
    | user_id        | PKEY      | ✅       | ✅          |
    | email          | TEXT      | ✅       |             |
    | enc_pw         | TEXT      | ✅       |             |
    | user_role      | ENUM      | ✅       |             |
    | last_logged_in | TIMESTAMP | ✅       |             |
    | created_at     | TIMESTAMP | ✅       |             |
    | updated_at     | TIMESTAMP | ✅       |             |
    
    Example:
        user_id = 1
        email = user@gmail.com
        enc_pw = 
        role = Default | Admin

#### Benchmarks:

    | Column Name                    | Datatype  | Not Null | Primary Key |
    |--------------------------------|-----------|----------|-------------|
    | benchmark_id                   | PKEY      | ✅       | ✅          |
    | name                           | TEXT      | ✅       |             |
    | description                    | TEXT      |          |             |
    | asset_allocation               | JSONB     | ✅       |             |
    | std_dev_pct                    | REAL      |          |             |
    | real_return_pct                | REAL      |          |             |
    | drawdown_yrs                   | INTEGER   |          |             |
    | rec_rebalance_threshold_pct    | INTEGER   |          |             |
    | is_deprecated                  | BOOLEAN   | ✅       |             |
    | user_id                        | INTEGER   | ✅       |             |
    | created_at                     | TIMESTAMP | ✅       |             |
    | updated_at                     | TIMESTAMP | ✅       |             |

    Example:
        benchmark_id = 1
        name = "Golden Butterfly"
        description: "A higher return portfolio based around the philosophies of the Permanent Portfolio"
        asset_allocation: {"TSM": 20, "DSCV": 20, "LTB": 20, "STB": 20, "GOLD": 20}
        std_dev_pct: 11.2 
        real_return_pct: 7.4
        drawdown_yrs: 3
        rec_rebalance_threshold_pct: 10
        is_deprecated: false
        user_id = 1

#### Snapshots

    | Column Name                | Datatype  | Not Null | Primary Key |
    |----------------------------|-----------|----------|-------------|
    | snap_id                    | PKEY      | ✅       | ✅          |
    | description                | TEXT      |          |             |
    | snap_date                  | DATE      | ✅       |             |
    | total                      | REAL      |          |             |
    | weighted_er_pct            | REAL      |          |             |
    | benchmark_id               | INTEGER   |          |             |
    | rebalance_threshold_pct    | INTEGER   |          |             |
    | value_order                | INTEGER[] |          |             |
    | user_id                    | INTEGER   | ✅       |             |
    | created_at                 | TIMESTAMP | ✅       |             |
    | updated_at                 | TIMESTAMP | ✅       |             |

    Example:
        snap_id = 1
        snap_date = 07/14/2024
        total = 100000.00
        weighted_er_pct = 0.125
        benchmark_id = 1
        rebalance_threshold_pct = 10
        value_order = [1, 3, 2, 5, 4]
        user_id = 1

#### SnapshotValues

    | Column Name    | Datatype  | Not Null | Primary Key |
    |----------------|-----------|----------|-------------|
    | snap_val_id    | PKEY      | ✅       | ✅          |
    | snap_id        | INTEGER   | ✅       |             |
    | account_id     | INTEGER   | ✅       |             |
    | holding_id     | INTEGER   | ✅       |             |
    | total          | REAL      | ✅       |             |
    | skip_rebalance | BOOLEAN   | ✅       |             |
    | user_id        | INTEGER   | ✅       |             |
    | created_at     | TIMESTAMP | ✅       |             |
    | updated_at     | TIMESTAMP | ✅       |             |

    Example:
        snap_val_id = 1
        snap_id = 1
        account_id = 1
        holding_id = 2 
        total = 5750.45
        skip_rebalance = false
        user_id = 1

#### Accounts:

    | Column Name    | Datatype  | Not Null | Primary Key |
    |----------------|-----------|----------|-------------|
    | account_id     | PKEY      | ✅       | ✅          |
    | name           | TEXT      | ✅       |             |
    | description    | TEXT      |          |             |
    | tax_shelter    | ENUM      | ✅       |             |
    | institution    | TEXT      | ✅       |             |
    | is_deprecated  | BOOLEAN   | ✅       |             |
    | user_id        | INTEGER   | ✅       |             |
    | created_at     | TIMESTAMP | ✅       |             |
    | updated_at     | TIMESTAMP | ✅       |             |

    Example:  
        account_id = 1
        name = FID-0904
        description = Fidelity investment account ending in 0904
        tax_shelter = Taxable | Traditional | Roth | HSA | 529
        institution = Fidelity
        is_deprecated = false
        user_id = 1
        
#### Holdings:

A generalized data type typically used to represent a mutual fund/ETF holding, individual stock, bond, or other asset 

    | Column Name    | Datatype  | Not Null | Primary Key |
    |----------------|-----------|----------|-------------|
    | holding_id     | PKEY      | ✅       | ✅          |
    | name           | TEXT      | ✅       |             |
    | ticker         | TEXT      |          |             |
    | asset_category | ENUM      | ✅       |             |
    | expense_ratio  | REAL      |          |             |
    | maturation_date| TEXT      |          |             |
    | interest_rate  | REAL      |          |             |
    | is_deprecated  | BOOLEAN   | ✅       |             |
    | user_id        | INTEGER   | ✅       |             |
    | created_at     | TIMESTAMP | ✅       |             |
    | updated_at     | TIMESTAMP | ✅       |             |
    
    Asset Categories (ENUM):
        CASH, BILLS, STB, ITB, LTB, COMMODITIES, GOLD, REITS, TSM,
        DLCB, DLCG, DLCV, DLCM, DMCB, DMCG, DMCV, DMCM, DSCG, DSCB, DSCV, DSCM,
        ILCB, ILCG, ILCV, ILCM, IMCB, IMCG, IMCV, IMCM, ISCB, ISCG, ISCV, ISCM,
        CRYPTO, OTHER

    Example:
        holding_id = 1
        name = Vanguard Total Stock Market Index
        ticker = VTSAX
        asset_category = TSM 
        expense_ratio_pct = 0.08
        is_deprecated = false
        user_id = 1   

    Example:
        holding_id = 2
        name = OTC Bond 05/2026
        asset_category = STB
        maturation_date = 05/01/2026
        interest_rate_pct = 4.5
        is_deprecated = false
        user_id = 1
