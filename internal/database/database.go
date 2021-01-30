package database

import (
	"github.com/TerexNik/hand-meal/internal/config"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var SqlxDB = createSqlxConnection()

func createSqlxConnection() *sqlx.DB {
	var connPool = createConnectionPool()
	// Then set up sqlx and return the created DB reference
	nativeDB := stdlib.OpenDBFromPool(connPool)
	return sqlx.NewDb(nativeDB, "pgx")
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
