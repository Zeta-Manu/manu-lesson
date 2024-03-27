package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Zeta-Manu/manu-lesson/internal/repositories"
)

var _ IVideoController = &VideoController{}

type IVideoController interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
}

type VideoController struct {
	logger *zap.Logger
	repo   *repositories.VideoRepository
}

func NewVideoController(logger *zap.Logger, repo *repositories.VideoRepository) *VideoController {
	return &VideoController{
		logger: logger,
		repo:   repo,
	}
}

func (vc *VideoController) Get(c *gin.Context) {
}

func (vc *VideoController) Post(c *gin.Context) {
}
