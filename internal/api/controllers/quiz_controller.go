package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Zeta-Manu/manu-lesson/internal/domain"
	"github.com/Zeta-Manu/manu-lesson/internal/repositories"
)

var _ IQuizController = &QuizController{}

type IQuizController interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
	List(c *gin.Context)
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

// @Summary Get a quiz by ID
// @Description Get a quiz by ID
// @Tags quiz
// @Accept json
// @Produce json
// @Param id path int true "Quiz ID"
// @Success 200
// @Router /quiz/{id} [get]
func (qc *QuizController) Get(c *gin.Context) {
	id := c.Param("id")
	quiz, err := qc.repo.GetQuizQuestion(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": quiz})
}

// @Summary Create a new quiz
// @Description Create a new quiz
// @Tags quiz
// @Accept json
// @Produce json
// @Param quiz body domain.Quiz true "Create quiz"
// @Success 201
// @Router /quiz [post]
func (qc *QuizController) Post(c *gin.Context) {
	var quiz domain.Quiz
	if err := c.BindJSON(&quiz); err != nil {
		qc.logger.Error("Bad Request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := qc.repo.PostQuizQuestion(quiz.Question, quiz.Answer)
	if err != nil {
		qc.logger.Error("Error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	qc.logger.Info("Create the quiz successfully")
	c.Status(http.StatusCreated)
}

// @Summary List all quizzes
// @Description List all quizzes
// @Tags quiz
// @Accept json
// @Produce json
// @Success 200
// @Router /quiz [get]
func (qc *QuizController) List(c *gin.Context) {
	quizes, err := qc.repo.GetAllQuestions()
	if err != nil {
		qc.logger.Error("Error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	qc.logger.Info("Get", zap.String(fmt.Sprintf("%d", len(quizes)), "quizes"))
	c.JSON(http.StatusOK, gin.H{"data": quizes})
}
