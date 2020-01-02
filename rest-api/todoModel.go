package rest_api

import (
	"net/http"
	"strconv"

	"../settings"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TodoModel struct {
	gorm.Model
	Title     string `json:"title"`
	Completed int    `json:"completed,omitempty"`
}

type FormattedTodoModel struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed,omitempty"`
}

func CreateTodoModel(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	model := TodoModel{
		Title:     c.PostForm("title"),
		Completed: completed,
	}
	db := settings.Database()
	db.Save(&model)
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Todo model was created successfully.",
		"resourceId": model.ID,
	})
}

func FetchTodoModels(c *gin.Context) {
	var models []TodoModel
	var _models []FormattedTodoModel
	db := settings.Database()
	db.Find(&models)
	if len(models) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo models not found.",
		})
		return
	}
	for _, item := range models {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_models = append(_models, FormattedTodoModel{
			ID:        item.ID,
			Title:     item.Title,
			Completed: completed,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   _models,
	})
}

func FetchTodoModel(c *gin.Context) {
	var model TodoModel
	modelId := c.Param("id")
	db := settings.Database()
	db.First(&model, modelId)
	if model.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo model #" + modelId + " not found.",
		})
		return
	}

	completed := false
	if model.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_model := FormattedTodoModel{
		ID:        model.ID,
		Title:     model.Title,
		Completed: completed,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   _model,
	})
}

func UpdateTodoModel(c *gin.Context) {
	var model TodoModel
	tTodoModelId := c.Param("id")
	db := settings.Database()
	db.First(&model, tTodoModelId)
	if model.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo model #" + tTodoModelId + " not found.",
		})
		return
	}
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&model).Update(
		"title", c.PostForm("title"),
		"completed", completed,
	)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo model #" + tTodoModelId + " was updated successfully.",
	})
}

func DeleteTodoModel(c *gin.Context) {
	var model TodoModel
	modelId := c.Param("id")
	db := settings.Database()
	db.First(&model, modelId)
	if model.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo model #" + modelId + " not found.",
		})
		return
	}
	db.Delete(&model)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo model #" + modelId + " was deleted successfully.",
	})
}
