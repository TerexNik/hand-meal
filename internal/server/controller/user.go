package controller

import (
	"encoding/json"
	"github.com/TerexNik/hand-meal/internal/database/repository"
	"net/http"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {
	users, err := repository.GetAllUsers()
	if err != nil {
		res.Write([]byte(err.Error()))
	}
	body, err := json.Marshal(users)
	if err != nil {
		res.Write([]byte(err.Error()))
	}
	res.Write(body)
}
