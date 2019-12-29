package rest_api

import (
	"../settings"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type TodoModel struct {
	gorm.Model
	Name string `json:"name"`
}

type FormattedTodoModel struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CreateTodoModel(c *gin.Context) {
	model := TodoModel{
		Name: c.PostForm("name"),
	}
	db := settings.Database()
	db.Save(&model)
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Todo Model created successfully.",
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
			"message": "Todo Model not found.",
		})
		return
	}
	for _, item := range models {
		_models = append(_models, FormattedTodoModel{
			ID:   item.ID,
			Name: item.Name,
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
			"message": "Todo Model #" + modelId + " not found.",
		})
		return
	}
	_model := FormattedTodoModel{
		ID:   model.ID,
		Name: model.Name,
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
			"message": "Todo Model #" + tTodoModelId + " not found.",
		})
		return
	}
	db.Model(&model).Update("name", c.PostForm("name"))
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo Model #" + tTodoModelId + " updated successfully.",
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
			"message": "Todo Model #" + modelId + " not found.",
		})
		return
	}
	db.Delete(&model)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo Model #" + modelId + " deleted successfully.",
	})
}
