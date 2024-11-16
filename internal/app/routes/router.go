package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/serikdev/go-todo/internal/app/handler"
)

func InitRoutes(r *gin.Engine) {

	r.POST("/tasks", handler.CreateTask)
	r.GET("/tasks", handler.GetTask)
	r.GET("/tasks/:id", handler.GetById)
	r.PUT("/tasks/:id", handler.UpdateTask)
	r.DELETE("/tasks/:id", handler.DeleteTask)
	r.PATCH("/tasks/:id/complete", handler.CompleteTask)

}
