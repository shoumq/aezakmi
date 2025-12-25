package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	AgoraAppID   string
	AgoraAppCert string
	DBUser       string
	DBPassword   string
	DBName       string
	DBHost       string
	DBPort       string

	TokenTTL int
}

func Load() *Config {
	var cfg Config

	cfg.AgoraAppID = getEnv("AGORA_APP_ID", "")
	cfg.AgoraAppCert = getEnv("AGORA_APP_CERT", "")

	cfg.DBUser = getEnv("POSTGRES_USER", "postgres")
	cfg.DBPassword = getEnv("POSTGRES_PASSWORD", "postgres")
	cfg.DBName = getEnv("POSTGRES_DB", "agora")
	cfg.DBHost = getEnv("POSTGRES_HOST", "postgres")
	cfg.DBPort = getEnv("POSTGRES_PORT", "5432")

	log.Println("qwerty", cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort)
	fmt.Println("qwerty", cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort)

	ttlSec, err := strconv.Atoi(getEnv("TOKEN_TTL", "3600"))
	if err != nil {
		log.Println("invalid TOKEN_TTL, using 3600")
		ttlSec = 3600
	}
	cfg.TokenTTL = ttlSec

	return &cfg
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
