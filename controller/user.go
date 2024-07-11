package controller

import (
	"log"
	"net/http"
	"simple/user/auth/service"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user service.UserDTO

	var err error

	err = c.ShouldBindJSON(&user)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Counld not parse user data"})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Counld not Insert user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}

func GetAllUsers(c *gin.Context) {

	users, err := service.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Counld not get users"})
		return
	}

	c.JSON(http.StatusOK, users)

}
