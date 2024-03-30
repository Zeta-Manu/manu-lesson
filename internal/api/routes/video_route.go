package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Zeta-Manu/manu-lesson/config"
	"github.com/Zeta-Manu/manu-lesson/internal/adapters/db"
	"github.com/Zeta-Manu/manu-lesson/internal/adapters/s3"
	"github.com/Zeta-Manu/manu-lesson/internal/api/controllers"
	"github.com/Zeta-Manu/manu-lesson/internal/repositories"
)

func InitVideoRoutes(router *gin.Engine, logger *zap.Logger, dbAdapter *db.Database, s3Adapter *s3.S3Adapter, cdnConfig config.CloudFrontConfig) {
	videoRepo := repositories.NewVideoRepository(dbAdapter, s3Adapter)
	videoController := controllers.NewVideoController(logger, videoRepo, &cdnConfig)
	quiz := router.Group("/api/video")
	{
		quiz.GET("/:id", videoController.Get)
		quiz.GET("/", videoController.List)
		quiz.POST("/", videoController.Post)
	}
}
