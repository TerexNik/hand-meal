package apiserver

import (
	"database/sql"
	"github.com/TerexNik/hand-meal/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	s := newServer(store, sessionStore)
	logger := logrus.New()
	logger.Info("server started on port: ", config.BindAddr)
	return http.ListenAndServe(config.BindAddr, s)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
