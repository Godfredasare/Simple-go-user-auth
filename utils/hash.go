package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(byte), err
}

func VerifyPassword(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
