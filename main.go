package main

import (
	controllers "Server/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	taskController := new(controllers.TaskController)
	loginController := new(controllers.LoginController)

	authM := loginController.AuthMiddleware()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router.POST("/login", loginController.Login)
	router.GET("/tasks", authM, taskController.GetAllTasks)
	router.POST("/addtasks", authM, taskController.AddTasks)

	router.Run("localhost:8080")
}
