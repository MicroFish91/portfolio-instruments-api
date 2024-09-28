package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	ShortRequestLimit int
	LongRequestLimit  int
}

var Env Config = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		Port:              getEnv("PORT", ":3000"),
		ShortRequestLimit: getInt(getEnv("SHORT_REQUEST_LIMIT", "99999")),
		LongRequestLimit:  getInt(getEnv("LONG_REQUEST_LIMIT", "99999")),
	}
}

func getEnv(key string, fallback string) string {
	if env, ok := os.LookupEnv(key); ok {
		return env
	}
	return fallback
}

func getInt(env string) int {
	e, err := strconv.Atoi(env)
	if err != nil {
		log.Fatal(err)
	}
	return e
}
