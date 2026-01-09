package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kareemhamed001/blog/internal/handlers"
)

func SetupAuthRoutes(router *gin.Engine, authHandler handlers.AuthHandler) {
	auth := router.Group("/auth")

	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
}
