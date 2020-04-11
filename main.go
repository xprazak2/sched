package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/xprazak2/sched/users"
	"github.com/xprazak2/sched/database"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&users.UserModel{})
}


func main() {

	connString := flag.String("db", "", "A connection string to postgres")
	flag.Parse()

	db := database.Connect(*connString)
	migrate(db)
  defer db.Close()

	router := gin.Default()

	// TODO configure origin from command line to fix the security issue
	router.Use(cors.Default())

	userRoutes := router.Group("/users")
	userRoutes.GET("/", users.UsersHandler)
	userRoutes.GET("/:id", users.UserHandler)
	userRoutes.POST("/", users.UserCreateHandler)
	userRoutes.PUT("/:id", users.UserUpdateHandler)
	userRoutes.DELETE("/:id", users.UserDeleteHandler)

	router.Run()
}
