package controller

import (
	"fmt"
	"log"
	"net/http"
	"simple/user/auth/service"
	"simple/user/auth/utils"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product service.ProductDTO
	err := c.ShouldBindJSON(&product)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cound not parse data"})
		return
	}

	errMessage := utils.Validation(product)
	if len(errMessage) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "errors": errMessage})
		return
	}

	userId := c.GetInt64("userId")
	product.User_id = userId

	err = product.InsertProduct()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cound not create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created Successfully"})

}

func GetAllProducts(c *gin.Context) {

	product, err := service.FindAllProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cound not Get all products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": product})

}

func GetOneProducts(c *gin.Context) {
	id := c.Param("id")

	exist, err := service.ProductExist(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Product does not exist"})
		return
	}

	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Product with id could not be found"})
		return
	}

	product, err := service.FindOne(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cound not Get product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var product service.ProductDTO
	err := c.ShouldBindJSON(&product)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cound not parse data"})
		return
	}

	productUserID, err := service.GetProductUserId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cound not get product user id"})
		return
	}

	userId := c.GetInt64("userId")
	if productUserID != userId {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorize user"})
		return
	}

	result, err := product.Update(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cound not update product"})
		return
	}

	if result == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated Successfully"})

}
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	productUserID, err := service.GetProductUserId(id)
	fmt.Println(productUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cound not get product user id"})
		return
	}

	userId := c.GetInt64("userId")

	if productUserID != userId {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorize user"})
		return
	}

	result, err := service.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cound not update product"})
		return
	}

	if result == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted Successfully"})

}

func GetProductsByUser(c *gin.Context) {
	id := c.Param("id")

	product, err := service.UserProducts(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cound not get product by user"})
		return
	}

	c.JSON(http.StatusOK, product)

}
