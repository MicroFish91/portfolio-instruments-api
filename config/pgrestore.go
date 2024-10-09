package config

import "github.com/joho/godotenv"

type PgRestoreConfig struct {
	Rest_DbHost     string
	Rest_DbPort     int
	Rest_DbName     string
	Rest_DbUser     string
	Rest_DbPassword string
	Rest_SourcePath string
}

func GetPgRestoreConfig() PgRestoreConfig {
	godotenv.Load(".env.pgrestore")

	return PgRestoreConfig{
		Rest_DbHost:     getEnv("REST_DB_HOST", "localhost"),
		Rest_DbPort:     getInt(getEnv("REST_DB_PORT", "5432")),
		Rest_DbName:     getEnv("REST_DB_NAME", "postgres"),
		Rest_DbUser:     getEnv("REST_DB_USER", ""),
		Rest_DbPassword: getEnv("REST_DB_PASSWORD", ""),
		Rest_SourcePath: getEnv("REST_SOURCE_PATH", ""),
	}
}
