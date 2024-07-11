package controller

import (
	"log"
	"net/http"
	"simple/user/auth/service"
	"simple/user/auth/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user service.LoginDTO

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Counld not parse user data"})
		return
	}

	errorMessage := utils.Validation(user)
	if len(errorMessage) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "errors": errorMessage})
		return
	}

	err = user.ValidatesCredentials()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email/password"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorize user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successfully", "token": token})

}
