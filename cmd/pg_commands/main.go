package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MicroFish91/portfolio-instruments-api/config"
	pg "github.com/habx/pg-commands"
)

func main() {
	cmd := os.Args[len(os.Args)-1]
	if cmd == "dump" {
		pgDump()
	} else if cmd == "restore" {
		pgRestore()
	}
}

func pgDump() {
	c := config.GetPgDumpConfig()

	dump, err := pg.NewDump(&pg.Postgres{
		Host:     c.Dump_DbHost,
		Port:     c.Dump_DbPort,
		DB:       c.Dump_DbName,
		Username: c.Dump_DbUser,
		Password: c.Dump_DbPassword,
	})
	if err != nil {
		log.Fatal(err)
	}

	dump.ResetOptions()
	dump.Options = []string{"--no-owner", "--no-privileges"}

	dumpExec := dump.Exec(pg.ExecOptions{
		StreamPrint: false,
	})

	fmt.Println("pg_dump", dumpExec.FullCommand)

	if dumpExec.Error != nil {
		fmt.Println(dumpExec.Error.Err)
		fmt.Println(dumpExec.Output)
	} else {
		fmt.Println("Dump success")
		fmt.Println(dumpExec.File)
	}
}

func pgRestore() {
	c := config.GetPgRestoreConfig()

	restore, err := pg.NewRestore(&pg.Postgres{
		Host:     c.Rest_DbHost,
		Port:     c.Rest_DbPort,
		DB:       c.Rest_DbName,
		Username: c.Rest_DbUser,
		Password: c.Rest_DbPassword,
	})
	if err != nil {
		log.Fatal(err)
	}

	restore.ResetOptions()
	restore.Options = []string{"--no-owner", "--no-privileges"}

	restoreExec := restore.Exec(c.Rest_SourcePath, pg.ExecOptions{
		StreamPrint: false,
	})

	fmt.Println("pg_restore", restoreExec.FullCommand)

	if restoreExec.Error != nil {
		fmt.Println(restoreExec.Error.Err)
		fmt.Println(restoreExec.Output)
	} else {
		fmt.Println("Restore success")
		fmt.Println(restoreExec.Output)
	}
}
