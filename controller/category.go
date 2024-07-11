package controller

import (
	"log"
	"net/http"
	"simple/user/auth/service"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category service.CategoryDTO

	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = category.InsertCategory()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not insert category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "category inserted successfully"})

}

func GetAllCategories(c *gin.Context) {
	category, err := service.FindAll()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could Find data"})
		return
	}

	c.JSON(http.StatusOK, category)

}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category service.CategoryDTO

	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	rows, err := category.UpdateOne(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not update data"})
		return
	}

	if rows == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not find category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "category updated successfully"})

}
