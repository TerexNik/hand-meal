package server

import (
	"github.com/TerexNik/hand-meal/internal/server/controller"
	"net/http"
)

func configureMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/recipes", controller.CheckAuthorize(controller.HandleRecipe))
	mux.Handle("/users", controller.CheckAuthorize(controller.GetUsers))
	return mux
}

func StartServer(serverAddr string) *http.Server {
	return &http.Server{
		Addr:    serverAddr,
		Handler: configureMux(),
	}
}
