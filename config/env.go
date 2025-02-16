package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	PublicHost            string
	Port                  string
	DBAddress             string
	MaxOpenConns          int64
	MaxIdleConns          int64
	MaxIdleTime           int64
	JWTExpirationInSecond int64
	JWTSecret             string
}

var Env = initConfigStaging()

func initConfigProduction() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "https://api.holyycan.com"),
		Port:       getEnv("PORT", "8080"),
		DBAddress: fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=%s",
			getEnv("DB_USER", "holycan"),
			getEnv("DB_PASSWORD", "ramaa212!"),
			getEnv("DB_HOST", "172.17.0.1:5432"),
			getEnv("DB_NAME", "wgca"),
			getEnv("SSL_MODE", "disable"),
		),
		MaxOpenConns: getEnvAsInt("DB_MAX_OPEN_CONNS", 30),
		MaxIdleConns: getEnvAsInt("DB_MAX_IDLE_CONNS", 30),
		MaxIdleTime:  getEnvAsInt("DB_MAX_IDLE_TIME", 5),
	}
}

func initConfigStaging() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBAddress: fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=%s",
			getEnv("DB_USER", "holycan"),
			getEnv("DB_PASSWORD", "ramaa212!"),
			getEnv("DB_HOST", "localhost:5432"),
			getEnv("DB_NAME", "smartparkingdb"),
			getEnv("SSL_MODE", "disable"),
		),
		MaxOpenConns:          getEnvAsInt("DB_MAX_OPEN_CONNS", 30),
		MaxIdleConns:          getEnvAsInt("DB_MAX_IDLE_CONNS", 30),
		MaxIdleTime:           getEnvAsInt("DB_MAX_IDLE_TIME", 5),
		JWTExpirationInSecond: getEnvAsInt("JWT_EXP", 3600*24),
		JWTSecret:             getEnv("JWT_SECRET", "TEST_SECRET"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
