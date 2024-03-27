package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Zeta-Manu/manu-lesson/internal/repositories"
)

var _ IQuizController = &QuizController{}

type IQuizController interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
}

type QuizController struct {
	logger *zap.Logger
	repo   *repositories.QuizRepository
}

func NewQuizController(logger *zap.Logger, repo *repositories.QuizRepository) *QuizController {
	return &QuizController{
		logger: logger,
		repo:   repo,
	}
}

func (qc *QuizController) Get(c *gin.Context) {
}

func (qc *QuizController) Post(c *gin.Context) {
}
