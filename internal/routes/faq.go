package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kareemhamed001/blog/internal/handlers"
	"github.com/kareemhamed001/blog/internal/middlewares"
	"github.com/kareemhamed001/blog/internal/types"
)

func SetupFaqRoutes(router *gin.Engine, faqHandler handlers.FAQHandler, jwtSecret string) {
	auth := router.Group("/api", middlewares.HasRole([]types.UserRole{types.RoleAdmin, types.RoleMerchant}, jwtSecret))

	faqCategories := auth.Group("/faqs")
	faqCategories.GET("/", faqHandler.GetAllFAQs)
	faqCategories.GET("/:id", faqHandler.GetFAQByID)
	faqCategories.POST("/", faqHandler.CreateFAQ)
	faqCategories.PUT("/:id", faqHandler.UpdateFAQ)
	faqCategories.DELETE("/:id", faqHandler.DeleteFAQ)
}
