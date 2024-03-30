package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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
	Update(c *gin.Context)
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
// @Failure 404
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
// @Param req body domain.Quiz true "Create quiz"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /quiz [post]
func (qc *QuizController) Post(c *gin.Context) {
	var req domain.Quiz
	if err := c.BindJSON(&req); err != nil {
		qc.logger.Error("Bad Request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := qc.repo.PostQuizQuestion(req.Question, req.Answer)
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
// @Failure 500
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

// @Summary Update the quiz
// @Description Update the existing quiz
// @Tags quiz
// @Accept json
// @Produce json
// @Param req body domain.QuizQuestion true "Answer and Question is optional"
// @Success 200 {string} string "success"
// @Failure 500 {object} string "Internal Server Error"
// @Router /quiz [put]
func (qc *QuizController) Update(c *gin.Context) {
	var req domain.QuizQuestion
	if err := c.BindJSON(&req); err != nil {
		qc.logger.Error("Bad Request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := qc.repo.UpdateQuizQuestion(strconv.Itoa(req.ID), &req.Question, &req.Answer)
	if err != nil {
		qc.logger.Error("Failed to update the quiz table", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
