package main

import (
	"./rest-api"
	"./settings"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
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
	db.AutoMigrate(&rest_api.TodoModel{})
	router := gin.Default()
	EnvGinMode := os.Getenv("GIN_MODE")
	gin.SetMode(EnvGinMode)
	// Serve the frontend
	router.Use(static.Serve("/", static.LocalFile("./react-app/dist", true)))
	router.GET("/", home)
	v1 := router.Group("/api/v1/")
	{
		v1.POST("todo/", rest_api.CreateTodoModel)
		v1.GET("todo/", rest_api.FetchTodoModels)
		v1.GET("todo/:id", rest_api.FetchTodoModel)
		v1.PUT("todo/:id", rest_api.UpdateTodoModel)
		v1.DELETE("todo/:id", rest_api.DeleteTodoModel)
	}
	router.Run()
}
