package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/serikdev/go-todo/internal/app/model"
	"github.com/serikdev/go-todo/internal/app/repository"
)

func CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.DB.Create(&task)
	c.JSON(http.StatusOK, task)
}

func GetTask(c *gin.Context) {
	var tasks []model.Task
	if err := repository.DB.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve tasks"})
		return
	}

	if len(tasks) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "tasks not found"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func GetById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
	}

	var task model.Task
	if err := repository.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found tasks "})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	var task model.Task
	id := c.Param("id")

	if err := repository.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var updateTask model.Task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Title = updateTask.Title
	task.Description = updateTask.Description
	task.DueData = updateTask.DueData
	task.OverData = updateTask.OverData
	task.Completed = updateTask.Completed

	repository.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := repository.DB.Delete(&model.Task{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.Status(http.StatusOK)
}

func CompleteTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Task
	if err := repository.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	task.Completed = true

	if err := repository.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to complete the task"})
	}
	c.JSON(http.StatusOK, task)
}
