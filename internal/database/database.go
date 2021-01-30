package database

import (
	"database/sql"
	"github.com/TerexNik/hand-meal/internal/config"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	log "github.com/sirupsen/logrus"
)

var SqlxDB = createSqlxConnection()

func createSqlxConnection() *sql.DB {
	var connPool = createConnectionPool()
	// Then set up sqlx and return the created DB reference
	return stdlib.OpenDBFromPool(connPool)
}

func createConnectionPool() *pgx.ConnPool {
	cfg := config.CreateDBConfig()
	connPool, err := pgx.NewConnPool(cfg)
	if err != nil {
		log.Errorf("%+v", cfg)
		log.Fatalf(err.Error())
	}
	return connPool
}
