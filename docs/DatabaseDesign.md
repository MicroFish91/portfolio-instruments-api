## 🗄 Database Design

- Using _'s in model props so as to correspond directly with the SQL database values.  This allows us to easily marshal values using `row.Scan()` in our store logic.

#### Users: 

    | Column Name    | Datatype  | Not Null | Primary Key |
    |----------------|-----------|----------|-------------|
    | user_id        | PKEY      | ✅       | ✅          |
    | email          | TEXT      | ✅       |             |
    | password_hash  | TEXT      | ✅       |             |
    | role           | ENUM      | ✅       |             |
    | created_at     | TIMESTAMP | ✅       |             |
    | updated_at     | TIMESTAMP | ✅       |             |
    
    Example:
        user_id = 1
        email = user@gmail.com
        password_hash = 
        role = User | Admin

#### Settings:

    | Column Name    | Datatype  | Not Null | Primary Key |
    |----------------|-----------|----------|-------------|
    | setting_id     | PKEY      | ✅       | ✅          |
    | reb_thresh_pct | INTEGER   | ✅       |             |
    | vp_thresh_pct  | INTEGER   | ✅       |             |
    | vp_enabled     | BOOLEAN   | ✅       |             |
    | user_id        | INTEGER   | ✅       |             |
    | benchmark_id   | INTEGER   |          |             |
    | created_at     | TIMESTAMP | ✅       |             |
    | updated_at     | TIMESTAMP | ✅       |             |

    Example:
        setting_id = 1
        reb_thresh_pct = 10
        vp_thresh_pct = 10
        vp_enabled = true
        user_id = 1
        benchmark_id = 5

#### Benchmarks:

    | Column Name    | Datatype  | Not Null | Primary Key |
    |----------------|-----------|----------|-------------|
    | benchmark_id   | PKEY      | ✅       | ✅          |
    | name           | TEXT      | ✅       |             |
    | description    | TEXT      |          |             |
    | std_dev_pct    | REAL      |          |             |
    | real_return_pct| REAL      |          |             |
    | drawdown_yrs   | INTEGER   |          |             |
    | is_deprecated  | BOOLEAN   | ✅       |             |
    | user_id        | INTEGER   | ✅       |             |
    | created_at     | TIMESTAMP | ✅       |             |
    | updated_at     | TIMESTAMP | ✅       |             |

    Example:
        benchmark_id = 1
        name = "Golden Butterfly"
        description: "A higher return portfolio based around the philosophies of the Permanent Portfolio"
        std_dev_pct: 11.2 
        real_return_pct: 7.4
        drawdown_yrs: 3
        is_deprecated: false
        user_id = 1

#### Snapshots

    | Column Name    | Datatype  | Not Null | Primary Key |
    |----------------|-----------|----------|-------------|
    | snap_id        | PKEY      | ✅       | ✅          |
    | description    | TEXT      |          |             |
    | snap_date      | DATE      | ✅       |             |
    | total          | REAL      | ✅       |             |
    | user_id        | INTEGER   | ✅       |             |
    | created_at     | TIMESTAMP | ✅       |             |
    | updated_at     | TIMESTAMP | ✅       |             |

    Example:
        snapshot_id = 1
        snapshot_date = 07/14/2024
        total = 100.00
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
        snapshot_id = 1
        account_id = 1
        holding_id = 2 
        total = 5750.45
        skip_rebalance = false // We want to rebalance by default
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
    
    Example:
        holding_id = 1
        name = Vanguard Total Stock Market Index
        ticker = VTSAX
        asset_category = TSM 
        expense_ratio = 0.08
        is_deprecated = false
        user_id = 1   

    Example:
        holding_id = 2
        name = OTC Bond 05/2026
        asset_category = STB
        maturation_date = 05/01/2026
        interest_rate = 4.5
        is_deprecated = false
        user_id = 1
