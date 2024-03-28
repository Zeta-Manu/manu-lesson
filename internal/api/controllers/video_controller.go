package controllers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Zeta-Manu/manu-lesson/internal/repositories"
)

var _ IVideoController = &VideoController{}

type IVideoController interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
	List(c *gin.Context)
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

// @Summary Get the video url
// @Description Get CloudFront url for specific id
// @Tags video
// @Produce json
// @Param id path string true "Video ID"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /video/{id} [get]
func (vc *VideoController) Get(c *gin.Context) {
	id := c.Param("id")

	video, err := vc.repo.GetVideo(id)
	if err != nil {
		vc.logger.Error("Failed to fetch the video", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": video})
}

// @Summary Upload a video
// @Description Upload a video to S3
// @Tags video
// @Produce json
// @Param file formData file true "Video file"
// @Param key formData string true "Unique key for the video"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /video [post]
func (vc *VideoController) Post(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to open file"})
		return
	}
	defer file.Close()

	// Read the file content into a byte slice
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	key := c.PostForm("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Key is required"})
		return
	}

	err = vc.repo.PostVideo(key, fileBytes)
	if err != nil {
		// Instead of sending two separate JSON responses, combine the error and data into one JSON object
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload video"})
		return
	}

	// If the upload is successful, you can return the data as part of the JSON response
	url := "mock-url" // This should be the actual URL of the uploaded file
	c.JSON(http.StatusOK, gin.H{"data": url})
}

// @Summary List out all videos
// @Description List out all videos in s3
// @Tags video
// @Produce json
// @Success 200
// @Failure 500
// @Router /video [get]
func (vc *VideoController) List(c *gin.Context) {
	videos, err := vc.repo.GetAllVideo()
	if err != nil {
		vc.logger.Error("Failed to fetch videos", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch videos"})
	}

	c.JSON(http.StatusOK, gin.H{"data": videos})
}
