package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strconv"
	"github.com/xprazak2/sched/database"
	"github.com/xprazak2/sched/errors"
)

func UsersHandler(ctx *gin.Context) {
	var users []UserModel
	db := database.Get()
	err := db.Find(&users).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("%v", err) })
	}
	ctx.JSON(http.StatusOK, gin.H{"users": ToResponse(users) })
}

func UserHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid id"})
	}
	db := database.Get()
	var user UserModel
	err = db.First(&user, id).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User with id %v not found", id)})
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user.ToResponse() })
}

func UserDeleteHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid id"})
	}
	db := database.Get()
	var user UserModel
	err = db.First(&user, id).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User with id %v not found", id)})
	}

	err = db.Delete(&user).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("%v", err) })
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user.ToResponse() })
}

func UserUpdateHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid id"})
		return
	}

	db := database.Get()
	var user UserModel
	err = db.First(&user, id).Error

	if &user != nil {
		validator := UpdateUserValidator(user)

		if err := validator.Bind(ctx); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"validationError": errors.NewValidationError(err)})
		} else {
			fmt.Sprintf("%+v", &validator.user)

			ctx.JSON(http.StatusOK, gin.H{"user": validator.user.ToResponse() })
		}
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User with id %v not found", id)})
	}
}

func UserCreateHandler(ctx *gin.Context) {
	validator := CreateUserValidator()

	if err := validator.Bind(ctx); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"validationError": errors.NewValidationError(err)})
	} else {
		db := database.Get()
		err := db.Create(&validator.user).Error

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("%v", err) })
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"user": validator.user.ToResponse() })
	}
}
