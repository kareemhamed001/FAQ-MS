package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kareemhamed001/blog/internal/handlers"
	"github.com/kareemhamed001/blog/internal/middlewares"
	"github.com/kareemhamed001/blog/internal/types"
)

func SetupFaqCategoriesRoutes(router *gin.Engine, faqCategoryHandler handlers.FAQCategoryHandler, jwtSecret string) {
	auth := router.Group("/api", middlewares.HasRole([]types.UserRole{types.RoleAdmin}, jwtSecret))

	faqCategories := auth.Group("/faq-categories")
	faqCategories.GET("/", faqCategoryHandler.GetAllCategories)
	faqCategories.GET("/:id", faqCategoryHandler.GetCategoryByID)
	faqCategories.POST("/", faqCategoryHandler.CreateCategory)
	faqCategories.PUT("/:id", faqCategoryHandler.UpdateCategory)
	faqCategories.DELETE("/:id", faqCategoryHandler.DeleteCategory)
}
