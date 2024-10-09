package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MicroFish91/portfolio-instruments-api/config"
	"github.com/MicroFish91/portfolio-instruments-api/db"
	"github.com/MicroFish91/portfolio-instruments-api/migrator"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	c := config.GetAppConfig()

	dbConfig := db.PostgresDbConfig{
		DbHost:     c.DbHost,
		DbPort:     c.DbPort,
		DbName:     c.DbName,
		DbUser:     c.DbUser,
		DbPassword: c.DbPassword,
		DbSslMode:  c.DbSslMode,
	}

	m, err := migrator.NewPostgresMigrator(
		db.GetDbConnectionString(dbConfig),
		"file://migrations", // migration source folder
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("successfully ran up migration")
	} else if cmd == "down" {
		if c.AppEnv == "production" {
			log.Fatal("blocking down migrations for prod")
		}

		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("successfully ran down migration")
	}
}
