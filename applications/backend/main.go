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

		router.GET("/users", controllers.FindUsers)
		router.POST("/users", controllers.StoreUser)
		router.GET("/users/:id", controllers.FindUserById)
		router.PUT("/users/:id", controllers.UpdateUser)
		router.DELETE("/users/:id", controllers.DeleteUser)

		router.Run(":5000")

}
