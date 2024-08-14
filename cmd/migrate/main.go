package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MicroFish91/portfolio-instruments-api/migrator"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	m, err := migrator.NewPostgresMigrator(
		"postgresql://localhost:5432/postgres?sslmode=disable", // connection string
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
		fmt.Println("successfully ran up migration")
	} else if cmd == "down" {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("successfully ran down migration")
	}
}
