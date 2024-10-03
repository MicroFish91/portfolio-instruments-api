package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv                   string
	Port                     string
	JwtSecret                string
	RequireVerification      bool
	UnauthorizedRequestLimit int
	ShortRequestLimit        int
	LongRequestLimit         int
	LogLevel                 int

	DbHost           string
	DbPort           string
	DbName           string
	DbUser           string
	DbPassword       string
	DbSslMode        string
	DbMaxConnections int
	DbMinConnections int
}

var Env Config = initConfig()

func initConfig() Config {
	godotenv.Load()

	var appEnv string = getEnv("APP_ENV", "development")
	if appEnv == "production" {
		fmt.Println("running in production")
		godotenv.Load(".env.prod")
	} else {
		fmt.Println("running in development")
		godotenv.Load(".env.dev")
	}

	return Config{
		AppEnv:                   appEnv,
		Port:                     getEnv("PORT", ":3000"),
		JwtSecret:                getEnv("JWT_SECRET", "not-so-secret-jwt-secret"),
		RequireVerification:      getBool(getEnv("REQUIRE_VERIFICATION", "true")),
		UnauthorizedRequestLimit: getInt(getEnv("UNAUTHORIZED_REQUEST_LIMIT", "99999")),
		ShortRequestLimit:        getInt(getEnv("SHORT_REQUEST_LIMIT", "99999")),
		LongRequestLimit:         getInt(getEnv("LONG_REQUEST_LIMIT", "99999")),
		LogLevel:                 getInt(getEnv("LOG_LEVEL", "0")),

		DbHost:           getEnv("DB_HOST", "localhost"),
		DbPort:           getEnv("DB_PORT", "5432"),
		DbName:           getEnv("DB_NAME", "postgres"),
		DbUser:           getEnv("DB_USER", ""),
		DbPassword:       getEnv("DB_PASSWORD", ""),
		DbSslMode:        getEnv("DB_SSL_MODE", "disable"),
		DbMaxConnections: getInt(getEnv("DB_MAX_CONNECTIONS", "4")),
		DbMinConnections: getInt(getEnv("DB_MIN_CONNECTIONS", "0")),
	}
}

func getEnv(key string, fallback string) string {
	if env, ok := os.LookupEnv(key); ok {
		return env
	}
	return fallback
}

func getBool(env string) bool {
	return env == "true"
}

func getInt(env string) int {
	i, err := strconv.Atoi(env)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
