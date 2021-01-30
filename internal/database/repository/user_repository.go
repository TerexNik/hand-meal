package repository

import (
	"github.com/TerexNik/hand-meal/internal/database"
	"github.com/TerexNik/hand-meal/internal/database/model"
)

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := database.SqlxDB.Select(&users, "select * from users")
	if err != nil {
		return nil, err
	}
	return users, nil
}
