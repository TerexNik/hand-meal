package config

import (
	"github.com/jackc/pgx"
	log "github.com/sirupsen/logrus"
	"os"
)

func ConfigureLogger(logLevel string) {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("Can't parse log level. LogLevel: %s", logLevel)
	}
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		PadLevelText:  true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(level)
}

type DBLogger struct {
}

func (dbLogger DBLogger) Log(level pgx.LogLevel, msg string, data map[string]interface{}) {
	log.Debug(" database ", msg, data)
}
