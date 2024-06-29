## 🗄 Database Design

#### Users: 

    | Column Name    | Datatype  | Not Null | Primary Key |
    |----------------|-----------|----------|-------------|
    | user_id        | UUID      | ✅       | ✅          |
    | email          | TEXT      | ✅       |             |
    | enc_password   | TEXT      | ✅       |             |
    | created_at     | TIMESTAMP | ✅       |             |
    | updated_at     | TIMESTAMP | ✅       |             |
    
#### Settings:
Todo

#### Snapshots:
Todo

#### Accounts:
Todo
    | Column Name    | Datatype  | Not Null | Primary Key |
    |----------------|-----------|----------|-------------|
    | account_id     | UUID      | ✅       | ✅          |
    | description    | TEXT      |          |             |
    | institution    | TEXT      | ✅       |             |
    | tax_shelter    | ENUM      | ✅       |             |
    | created_at     | TIMESTAMP | ✅       |             |
    | updated_at     | TIMESTAMP | ✅       |             |
    
### Holdings:
Todo 