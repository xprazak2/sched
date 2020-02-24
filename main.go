package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/xprazak2/sched/users"
	"github.com/xprazak2/sched/database"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&users.UserModel{})
}

// func initDb() {
// 	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gorm password=postgres")
// 	if err != nil {
// 		fmt.Println("db err: ", err)
// 	}
// 	return db
// }

func main() {

	db := database.Connect()
	migrate(db)
  defer db.Close()

	router := gin.Default()

	// router.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	userRoutes := router.Group("/users")
	userRoutes.GET("/", users.UsersHandler)
	userRoutes.GET("/:id", users.UserHandler)
	userRoutes.POST("/", users.UserCreateHandler)
	userRoutes.PUT("/:id", users.UserUpdateHandler)
	userRoutes.DELETE("/:id", users.UserDeleteHandler)

	router.Run()
}
