package main

import (
	controllers "Server/controllers"
	sseSever "Server/sse"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()

	taskController := new(controllers.TaskController)
	loginController := new(controllers.LoginController)
	userController := new(controllers.UserController)

	authM := loginController.AuthMiddleware()

	corsM := CORSMiddleware()
	router.Use(corsM)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router.POST("/login", loginController.Login)
	router.GET("/tasks", authM, taskController.GetAllTasks)
	router.GET("/tasks/:id", authM, taskController.GetTask)

	router.POST("/addtasks", authM, taskController.AddTasks)
	router.GET("/getuser/:id", authM, userController.GetUser)

	s := sseSever.SseServer

	router.GET("/events/:channel", func(c *gin.Context) {
		s.ServeHTTP(c.Writer, c.Request)
	})

	router.Run("localhost:8080")
}
