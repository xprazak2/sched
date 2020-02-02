package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UsersHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"users": Users })
}