package model

type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"  validate:"required,min=3,max=32"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=3"`
	Created_at string `json:"created_at" binding:"required"`
	Updated_at string `json:"updated_at" binding:"required"`
}
