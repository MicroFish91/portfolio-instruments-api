# Migrations

Guide to running db migrations (entrypoint: `cmd/migrate/`).

### Installation 

Follow the instructions to install the migrate CLI [here](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md).

### Examples of Code Setup

https://pkg.go.dev/github.com/golang-migrate/migrate/v4#section-readme

```golang
import (
    "database/sql"
    _ "github.com/lib/pq"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    db, err := sql.Open("postgres", "postgres://localhost:5432/database?sslmode=enable")
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    m, err := migrate.NewWithDatabaseInstance(
        "file:///migrations",
        "postgres", driver)
    m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
}
```

### Examples of Running

- `make migration create-user-table` (Scaffolds a .sql migration file)
- `make migrate-up`
- `make migrate-down`

### Issues

Had to use the [pg](https://github.com/lib/pq) database driver for migrations specifically due to a legacy issue where migrate hasn't been able to fully support pgx driver yet. pg is no longer actively maintained and the creators have explicitly said to go use pgx instead, which is why I'm using pgx in the actual api implementation. pg will only be used temporarily here for migrates only.