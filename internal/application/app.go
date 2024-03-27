package application

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"github.com/Zeta-Manu/manu-lesson/config"
	docs "github.com/Zeta-Manu/manu-lesson/docs"
	"github.com/Zeta-Manu/manu-lesson/internal/adapters/db"
	"github.com/Zeta-Manu/manu-lesson/internal/adapters/s3"
	"github.com/Zeta-Manu/manu-lesson/internal/api/routes"
)

func NewApplication(cfg config.AppConfig) {
	router := gin.Default()

	logger, _ := zap.NewProduction()
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "healthy"})
	})

	dbAdapter, err := db.InitializeDatabase(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	s3Adapter, err := s3.NewS3Adapter(cfg.AWS.AccessKey, cfg.AWS.SecretAccessKey, cfg.S3.BucketName, cfg.S3.Region)
	if err != nil {
		log.Fatalf("Failed to connect to S3: %v", err)
	}

	docs.SwaggerInfo.BasePath = "/api"

	routes.InitQuizRoutes(router, logger, dbAdapter)
	routes.InitVideoRoutes(router, logger, dbAdapter, s3Adapter)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	startServer(router, logger)
}

func startServer(handler http.Handler, logger *zap.Logger) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to Start Server", zap.Error(err))
		}
	}()
	logger.Info("Starting Server ...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", zap.Error(err))
	}

	select {
	case <-ctx.Done():
		logger.Info("timeout of 5 seconds.")
	}
	logger.Info("Server exiting")
}
