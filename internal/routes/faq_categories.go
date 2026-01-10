package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kareemhamed001/faq/internal/handlers"
	"github.com/kareemhamed001/faq/internal/middlewares"
	"github.com/kareemhamed001/faq/internal/types"
)

func SetupFaqCategoriesRoutes(router *gin.Engine, faqCategoryHandler handlers.FAQCategoryHandler, jwtSecret string) {
	auth := router.Group("/api")

	faqCategories := auth.Group("/faq-categories")
	faqCategories.GET("/", middlewares.HasRole([]types.UserRole{types.RoleAdmin, types.RoleMerchant}, jwtSecret), faqCategoryHandler.GetAllCategories)
	faqCategories.GET("/:id", middlewares.HasRole([]types.UserRole{types.RoleAdmin}, jwtSecret), faqCategoryHandler.GetCategoryByID)
	faqCategories.POST("/", middlewares.HasRole([]types.UserRole{types.RoleAdmin}, jwtSecret), faqCategoryHandler.CreateCategory)
	faqCategories.PUT("/:id", middlewares.HasRole([]types.UserRole{types.RoleAdmin}, jwtSecret), faqCategoryHandler.UpdateCategory)
	faqCategories.DELETE("/:id", middlewares.HasRole([]types.UserRole{types.RoleAdmin}, jwtSecret), faqCategoryHandler.DeleteCategory)
}
