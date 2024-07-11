package service

import (
	"errors"
	"log"
	"simple/user/auth/database"
	"simple/user/auth/utils"
)

type LoginDTO struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3"`
}

func (user *LoginDTO) ValidatesCredentials() error {
	query := `SELECT id, password, email FROM users WHERE email = $1`

	var retrievedPassword string
	err := database.DB.QueryRow(query, user.Email).Scan(&user.ID, &retrievedPassword, &user.Email)
	if err != nil {
		log.Println("Error", err)
		return errors.New("credentials invalid")
	}

	password := utils.VerifyPassword(user.Password, retrievedPassword)
	if !password {
		return errors.New("credentials invalid")
	}

	return nil
}
