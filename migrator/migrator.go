package migrator

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func NewPostgresMigrator(connStr string, sourceUrl string) (*migrate.Migrate, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return migrate.NewWithDatabaseInstance(
		sourceUrl,
		"postgres",
		driver,
	)
}
