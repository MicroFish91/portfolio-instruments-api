## 🗄 Database Design

- Using _'s in model props so as to correspond directly with the SQL database values.  This allows us to easily marshal values using `row.Scan()` in our store logic.

#### Users: 

    | Column Name    | Datatype  | Not Null | Primary Key |
    |----------------|-----------|----------|-------------|
    | user_id        | UUID      | ✅       | ✅          |
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
    | setting_id     | UUID      | ✅       | ✅          |
    | reb_threshold  | DECIMAL   | ✅       |             |
    | vp_threshold   | DECIMAL   | ✅       |             |
    | created_at     | TIMESTAMP | ✅       |             |
    | updated_at     | TIMESTAMP | ✅       |             |

    Example:
        setting_id = 1
        reb_threshold = 0.15
        vp_threshold = 0.1

#### Benchmarks:
Todo
Similar idea to holdings below.  Need global and user-defined benchmark support.

#### Snapshots
Todo

#### SnapshotValues
Todo 

#### Accounts:

    | Column Name    | Datatype  | Not Null | Primary Key |
    |----------------|-----------|----------|-------------|
    | account_id     | UUID      | ✅       | ✅          |
    | name           | TEXT      | ✅       |             |
    | description    | TEXT      |          |             |
    | shelter_type   | ENUM      | ✅       |             |
    | institution    | TEXT      | ✅       |             |
    | user_id        | INTEGER   | ✅       |             |
    | created_at     | TIMESTAMP | ✅       |             |
    | updated_at     | TIMESTAMP | ✅       |             |
    | deleted_at     | TIMESTAMP |          |             |

    Example:  
        account_id = 1
        name = FID-0904
        description = Fidelity investment account ending in 0904
        <!-- account_type = savings -->
        tax_shelter = Taxable | Traditional | Roth | HSA | 529
        institution = Fidelity
        user_id = 1
        
#### Holdings:
When querying, users should be able to see the global-defined holdings as well as the user-defined holdings.

    | Column Name    | Datatype  | Not Null | Primary Key |
    |----------------|-----------|----------|-------------|
    | holding_id     | UUID      | ✅       | ✅          |
    | name           | TEXT      |          |             |
    | ticker         | TEXT      | ✅       |             |
    | category       | ENUM      | ✅       |             |
    | expense_ratio  | DECIMAL   | ✅       |             |
    | user_id        | INTEGER   | ✅       |             |
    | created_at     | TIMESTAMP | ✅       |             |
    | updated_at     | TIMESTAMP | ✅       |             |
    | deleted_at     | TIMESTAMP |          |             |
    
    Example:
        holding_id = 1
        name = Vanguard Total Stock Market Index
        ticker = VTSAX
        category = TSM,DLCB 
        expense_ratio = 0.08
        user_id = 1   
