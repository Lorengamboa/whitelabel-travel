package main

import (
	"encoding/hex"
	"flag"
	"log"
	"os"
	"strconv"
	"time"
)

func updateConfigWithEnvVariables() (*config, error) {
	// Load environment variables from `.env` file
	var cfg config
	debug, err := strconv.ParseBool("true")
	if err != nil {
		log.Fatal(err)
	}
	flag.BoolVar(&cfg.debug, "debug", debug, "Debug (true|false)")
	flag.Parse()

	// if cfg.debug {
	// 	err = godotenv.Load()
	// 	if err != nil {
	// 		log.Fatal("Error loading .env file: ", err)
	// 	}

	// }

	maxOpenConnsStr := os.Getenv("DB_MAX_OPEN_CONNS")
	maxOpenConns, err := strconv.Atoi(maxOpenConnsStr)
	if err != nil {
		log.Fatal(err)
	}
	maxIdleConnsStr := os.Getenv("DB_MAX_IDLE_CONNS")
	maxIdleConns, err := strconv.Atoi(maxIdleConnsStr)
	if err != nil {
		log.Fatal(err)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	// Basic config
	flag.IntVar(&cfg.port, "port", port, "API server port")
	// Database config
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DATABASE_URL"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", maxOpenConns, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", maxIdleConns, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime,
		"db-max-idle-time",
		os.Getenv("DB_MAX_IDLE_TIME"),
		"PostgreSQL max connection idle time",
	)

	// Secret
	flag.StringVar(&cfg.secret.HMC, "secret-key", os.Getenv("HMC_SECRET_KEY"), "HMC Secret Key")

	flag.Parse()

	secretKey, err := hex.DecodeString(cfg.secret.HMC)
	if err != nil {
		return nil, err
	}
	cfg.secret.secretKey = secretKey
	sessionDuration, err := time.ParseDuration(os.Getenv("SESSION_EXPIRATION"))
	if err != nil {
		return nil, err
	}
	cfg.secret.sessionExpiration = sessionDuration

	// Token Expiration
	tokexpirationStr := os.Getenv("TOKEN_EXPIRATION")
	duration, err := time.ParseDuration(tokexpirationStr)
	if err != nil {
		return nil, err
	}
	cfg.tokenExpiration.durationString = tokexpirationStr
	cfg.tokenExpiration.duration = duration

	return &cfg, nil
}
