package main

import (
		"github.com/gin-gonic/gin"
		"cloudops/backend/controllers"
		"cloudops/backend/models"
)

func main() {
		router := gin.Default()

		models.ConnectDatabase()

		router.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})

		router.GET("/tasks", controllers.FindTasks)
		router.POST("/tasks", controllers.StoreTask)
		router.GET("/tasks/:id", controllers.FindTaskById)
		router.PUT("/tasks/:id", controllers.UpdateTask)
		router.DELETE("/tasks/:id", controllers.DeleteTask)

		router.Run(":5000")

}
