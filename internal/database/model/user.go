package model

type User struct {
	Id                int    `json:"id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"-" db:"encrypted_password"`
	FIO               string `json:"fio" db:"fio"`
}
