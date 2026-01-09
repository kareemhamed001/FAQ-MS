package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kareemhamed001/blog/internal/handlers"
)

func SetupStoreRoutes(router *gin.Engine, storeHandler handlers.StoreHandler) {
	stores := router.Group("/api/stores")
	stores.GET("/", storeHandler.ListStores)
	stores.GET("/:id", storeHandler.GetStore)
}
