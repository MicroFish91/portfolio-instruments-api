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
	dbConfig := db.PostgresDbConfig{
		DbHost:     config.Env.DbHost,
		DbPort:     config.Env.DbPort,
		DbName:     config.Env.DbName,
		DbUser:     config.Env.DbUser,
		DbPassword: config.Env.DbPassword,
		DbSslMode:  config.Env.DbSslMode,
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
		if config.Env.AppEnv == "production" {
			log.Fatal("blocking down migrations for prod")
		}

		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("successfully ran down migration")
	}
}
