package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/kareemhamed001/blog/internal/DB"
	"github.com/kareemhamed001/blog/internal/config"
	"github.com/kareemhamed001/blog/internal/handlers"
	"github.com/kareemhamed001/blog/internal/logger"
	"github.com/kareemhamed001/blog/internal/middlewares"
	"github.com/kareemhamed001/blog/internal/routes"
	"github.com/kareemhamed001/blog/internal/services"
)

func main() {

	config := config.NewConfig()
	logr := logger.New(config.AppEnv)
	defer logr.Sync()

	if config.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := db.InitializeDB(config.DBDriver, config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		logr.Errorw("failed to connect to database", "error", err)
		log.Fatal("Failed to connect to database:", err)
	}

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middlewares.SetUserData(config.JWTPrivateKey))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	// Auth Routes
	authService := services.NewAuthService(db, config.JWTPrivateKey)
	authHandler := handlers.NewAuthHandler(*authService)

	routes.SetupAuthRoutes(router, *authHandler)

	// FAQ Category Routes
	faqCategoryService := services.NewFAQCategoryService(db)
	faqCategoryHandler := handlers.NewFAQCategoryHandler(*faqCategoryService)

	routes.SetupFaqCategoriesRoutes(router, *faqCategoryHandler, config.JWTPrivateKey)

	// FAQ  Routes
	faqService := services.NewFAQService(db)
	storeService := services.NewStoreService(db)
	faqHandler := handlers.NewFAQHandler(*faqService, *storeService)
	storeHandler := handlers.NewStoreHandler(*storeService)

	routes.SetupFaqRoutes(router, *faqHandler, config.JWTPrivateKey)
	routes.SetupStoreRoutes(router, *storeHandler)

	logr.Infow("starting server", "port", config.AppPort)
	router.Run(":" + strconv.Itoa(config.AppPort))
}
