package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateToken(userId int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(secretKey)
}

func VerifyToken(token string) (int64, error) {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secretKey, nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	if !parseToken.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userId := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("userId not found in token claims")
	}

	fmt.Println(userId)

	return int64(userId), nil
}
