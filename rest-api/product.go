package rest_api

import (
	"../settings"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Product struct {
	gorm.Model
	Name string `json:"name"`
}

type FormattedProduct struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CreateProduct(c *gin.Context) {
	product := Product{
		Name: c.PostForm("name"),
	}
	db := settings.Database()
	db.Save(&product)
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Product item created successfully.",
		"resourceId": product.ID,
	})
}

func FetchProducts(c *gin.Context) {
	var products []Product
	var _products []FormattedProduct
	db := settings.Database()
	db.Find(&products)
	if len(products) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Product not found",
		})
		return
	}
	for _, item := range products {
		_products = append(_products, FormattedProduct{
			ID:   item.ID,
			Name: item.Name,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   _products,
	})
}

func FetchProduct(c *gin.Context) {
	var product Product
	productId := c.Param("id")
	db := settings.Database()
	db.First(&product, productId)
	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Product " + productId + " not found",
		})
		return
	}
	_product := FormattedProduct{
		ID:   product.ID,
		Name: product.Name,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   _product,
	})
}

func UpdateProduct(c *gin.Context) {
	var product Product
	tProductId := c.Param("id")
	db := settings.Database()
	db.First(&product, tProductId)
	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Product " + tProductId + " not found",
		})
		return
	}
	db.Model(&product).Update("name", c.PostForm("name"))
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Product " + tProductId + " updated successfully",
	})
}

func DeleteProduct(c *gin.Context) {
	var product Product
	productId := c.Param("id")
	db := settings.Database()
	db.First(&product, productId)
	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Product " + productId + " not found",
		})
		return
	}
	db.Delete(&product)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Product " + productId + " deleted successfully",
	})
}
