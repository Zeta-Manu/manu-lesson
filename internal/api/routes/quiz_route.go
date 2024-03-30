package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Zeta-Manu/manu-lesson/internal/adapters/db"
	"github.com/Zeta-Manu/manu-lesson/internal/api/controllers"
	"github.com/Zeta-Manu/manu-lesson/internal/repositories"
)

func InitQuizRoutes(router *gin.Engine, logger *zap.Logger, dbAdapter *db.Database) {
	quizRepo := repositories.NewQuizRepository(dbAdapter)
	quizController := controllers.NewQuizController(logger, quizRepo)
	quiz := router.Group("/api/quiz")
	{
		quiz.GET("/:id", quizController.Get)
		quiz.GET("/", quizController.List)
		quiz.POST("/", quizController.Post)
		quiz.PUT("/", quizController.Update)
	}
}
