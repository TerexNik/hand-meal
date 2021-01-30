package repository

import (
	"github.com/TerexNik/hand-meal/internal/database"
	"github.com/TerexNik/hand-meal/internal/database/model"
)

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	_, err := database.SqlxDB.Query("select * from users", users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
