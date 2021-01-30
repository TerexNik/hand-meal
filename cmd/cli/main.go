package main

import (
	"github.com/TerexNik/hand-meal/internal/config"
	server2 "github.com/TerexNik/hand-meal/internal/server"
	"github.com/goanywhere/env"
	log "github.com/sirupsen/logrus"
)

func main() {
	env.Load(".env")
	config.ConfigureLogger(env.String("LOG_LEVEL", "debug"))

	s := server2.StartServer(env.String("SERVER_PORT", ":8080"))
	defer s.Close()
	log.Info("Server starts")
	s.ListenAndServe()
}
