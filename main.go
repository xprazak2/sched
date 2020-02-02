package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xprazak2/sched/users"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userRoutes := router.Group("/users")
	userRoutes.GET("/", users.UsersHandler)

	router.Run()
}
