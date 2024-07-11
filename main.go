package main

import (
	"log"
	"simple/user/auth/config"
	"simple/user/auth/database"
	"simple/user/auth/router"
	"simple/user/auth/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	database.InitDB()

	utils.InitializeValidatorUniversalTranslator()

	server := gin.Default()

	router.Router(server)

	server.Run(":4000")
}
