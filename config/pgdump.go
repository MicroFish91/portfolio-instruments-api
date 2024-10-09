package config

import "github.com/joho/godotenv"

type PgDumpConfig struct {
	Dump_DbHost     string
	Dump_DbPort     int
	Dump_DbName     string
	Dump_DbUser     string
	Dump_DbPassword string
}

func GetPgDumpConfig() PgDumpConfig {
	godotenv.Load(".env.pgdump")

	return PgDumpConfig{
		Dump_DbHost:     getEnv("DUMP_DB_HOST", "localhost"),
		Dump_DbPort:     getInt(getEnv("DUMP_DB_PORT", "5432")),
		Dump_DbName:     getEnv("DUMP_DB_NAME", "postgres"),
		Dump_DbUser:     getEnv("DUMP_DB_USER", ""),
		Dump_DbPassword: getEnv("DUMP_DB_PASSWORD", ""),
	}
}
