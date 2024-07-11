package service

import (
	"log"
	"simple/user/auth/database"
	"simple/user/auth/model"
	"simple/user/auth/utils"
)

type UserDTO struct {
	Username   string `json:"username" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=3"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

func (user UserDTO) Save() error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("Error", err)
		return err
	}

	_, err = database.DB.Exec(query, user.Username, user.Email, hashPassword)
	if err != nil {
		log.Println("Error inserting user")
		return err
	}
	return nil
}

func FindAllUsers() ([]model.User, error) {
	query := `SELECT id, username, email, password, created_at, updated_at FROM users`

	rows, err := database.DB.Query(query)
	if err != nil {
		log.Println("Error Getting all users user", err)
		return []model.User{}, err
	}

	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
		if err != nil {
			log.Println("Error Getting all users user", err)
			return []model.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}
