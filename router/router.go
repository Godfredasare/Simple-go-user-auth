package router

import (
	"simple/user/auth/controller"
	"simple/user/auth/middleware"

	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {
	server.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Jumia.com"})
	})

	server.POST("/api/category", controller.CreateCategory)
	server.GET("/api/category", controller.GetAllCategories)
	server.PUT("/api/category/:id", controller.UpdateCategory)

	server.GET("/api/product", controller.GetAllProducts)
	server.GET("/api/product/:id", controller.GetOneProducts)
	server.GET("/api/product/user/:id", controller.GetProductsByUser)

	Authenticated := server.Group("/api/")
	Authenticated.Use(middleware.AuthMiddleware)

	Authenticated.POST("product", controller.CreateProduct)
	Authenticated.PUT("product/:id", controller.UpdateProduct)
	Authenticated.DELETE("product/:id", controller.DeleteProduct)


	server.POST("/api/user", controller.CreateUser)
	server.GET("/api/user", controller.GetAllUsers)
	server.POST("/api/login", controller.Login)

}
