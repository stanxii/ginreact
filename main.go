package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	rest_api "./rest-api"
	"./settings"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db = settings.Database()
	db.AutoMigrate(&rest_api.TodoModel{})
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "REST API with Gin",
	})
}

func main() {
	router := gin.Default()
	EnvGinMode := os.Getenv("GIN_MODE")
	gin.SetMode(EnvGinMode)
	// Serve the frontend
	router.Use(static.Serve("/", static.LocalFile("./react-app/dist", true)))
	router.Use(CORSMiddleware())
	router.GET("/", home)
	v1 := router.Group("/api/v1/")
	{
		v1.POST("todo/", rest_api.CreateTodoModel)
		v1.GET("todo/", rest_api.FetchTodoModels)
		v1.GET("todo/:id", rest_api.FetchTodoModel)
		v1.PUT("todo/:id", rest_api.UpdateTodoModel)
		v1.DELETE("todo/:id", rest_api.DeleteTodoModel)
	}
	router.Run(":8000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Accept", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Access-Control-Allow-Headers, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, DELETE, GET, PUT")
		c.Next()
	}
}
