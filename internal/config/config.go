package config

import (
	"github.com/goanywhere/env"
	"github.com/jackc/pgx"
	log "github.com/sirupsen/logrus"
	"time"
)

type Config struct {
}

func CreateDBConfig() pgx.ConnPoolConfig {
	timeout, err := time.ParseDuration(env.String("DATABASE_CONNECTION_TIMEOUT", "2m"))
	if err != nil {
		log.Fatalf("Error when parse DATABASE_CONNECTION_TIMEOUT env")
	}
	return pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     env.String("DATABASE_HOST", "localhost"),
			Port:     uint16(env.Uint("DATABASE_PORT", 5432)),
			Database: env.String("DATABASE_NAME", "hm"),
			User:     env.String("DATABASE_USERNAME", "handmeal"),
			Password: env.String("DATABASE_PASSWORD", "qaz123"),
			Logger:   DBLogger{},
		},
		MaxConnections: env.Int("DATABASE_CONNECTION_LIMIT", 20),
		AfterConnect:   nil,
		AcquireTimeout: timeout,
	}
}
