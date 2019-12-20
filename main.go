package main

import (
	"./rest-api"
	"./settings"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "REST API with Gin",
	})
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db := settings.Database()
	db.AutoMigrate(&rest_api.Product{})
	router := gin.Default()
	EnvGinMode := os.Getenv("GIN_MODE")
	gin.SetMode(EnvGinMode)
	router.GET("/", home)
	router.LoadHTMLGlob("templates/*")
	v1 := router.Group("/api/v1/")
	{
		v1.POST("product/", rest_api.CreateProduct)
		v1.GET("product/", rest_api.FetchProducts)
		v1.GET("product/:id", rest_api.FetchProduct)
		v1.PUT("product/:id", rest_api.UpdateProduct)
		v1.DELETE("product/:id", rest_api.DeleteProduct)
	}
	router.Run()
}
