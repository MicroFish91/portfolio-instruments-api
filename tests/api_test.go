package tests

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	pg "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestApi(t *testing.T) {
	ctx := context.Background()
	pgc, err := pg.Run(ctx,
		"postgres:16-alpine",
		pg.WithDatabase("pidb"),
		pg.WithUsername("piuser"),
		pg.WithPassword("pipassword"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(time.Second*5),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	connStr, err := pgc.ConnectionString(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(connStr)

	// connect golang migrate
	// run migrate up

	// setup app
	// run test routes by passing test app

	// app need a channel for closing so it can close and be triggered after finishing tests?

	// Each parallel test can use a unique user id

	t.Cleanup(func() {
		fmt.Println("----cleaning up postgres test container...----")
		if err := pgc.Terminate(ctx); err != nil {
			fmt.Printf("failed to terminate postgres test container: %s", err.Error())
		}
	})
}
