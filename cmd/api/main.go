package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/kareemhamed001/faq/internal/DB"
	"github.com/kareemhamed001/faq/internal/config"
	"github.com/kareemhamed001/faq/internal/handlers"
	"github.com/kareemhamed001/faq/internal/logger"
	"github.com/kareemhamed001/faq/internal/middlewares"
	"github.com/kareemhamed001/faq/internal/routes"
	"github.com/kareemhamed001/faq/internal/services"
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

	// Allow frontend origins to call the API during development.
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173", "http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept-Language"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
