package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDbConfig struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	DbSslMode  string
	DbUrl      string

	DbMaxConnections int
	DbMinConnections int
}

func NewPostgresStorage(dbConfig PostgresDbConfig) (*pgxpool.Pool, error) {
	var connStr string
	if dbConfig.DbUrl != "" {
		connStr = dbConfig.DbUrl
	} else {
		var dbLogin string
		if dbConfig.DbUser != "" && dbConfig.DbPassword != "" {
			dbLogin = fmt.Sprintf("%s:%s@", dbConfig.DbUser, dbConfig.DbPassword)
		}
		connStr = fmt.Sprintf("postgresql://%s%s:%s/%s?sslmode=%s", dbLogin, dbConfig.DbHost, dbConfig.DbPort, dbConfig.DbName, dbConfig.DbSslMode)
	}

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal(err)
	}

	config.MinConns = int32(dbConfig.DbMinConnections)
	config.MaxConns = int32(dbConfig.DbMaxConnections)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	return pgxpool.NewWithConfig(ctx, config)
}
