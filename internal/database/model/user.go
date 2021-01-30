package model

type User struct {
	Id                int    `json:"id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"encrypted_password"`
}
