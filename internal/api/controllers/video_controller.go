package controllers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Zeta-Manu/manu-lesson/config"
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
	cdn    string
}

func NewVideoController(logger *zap.Logger, repo *repositories.VideoRepository, cloudFront *config.CloudFrontConfig) *VideoController {
	return &VideoController{
		logger: logger,
		repo:   repo,
		cdn:    cloudFront.Domain,
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
// @Param handsign formData string true "Handsign contain in the video"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /video [post]
func (vc *VideoController) Post(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		vc.logger.Error("Failed to retrived file", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
		return
	}

	key := c.PostForm("key")
	if key == "" {
		vc.logger.Error("Failed to retrived key", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Key is required"})
		return
	}

	handsign := c.PostForm("handsign")
	if handsign == "" {
		vc.logger.Error("Failed to retrived handsign", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Handsign is required"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		vc.logger.Error("Cannot open the file", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to open file"})
		return
	}
	defer file.Close()

	// Read the file content into a byte slice
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		vc.logger.Error("Cannot read the file", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	const (
		FOLDERNAME = "lesson"
	)

	err = vc.repo.PostVideo(FOLDERNAME, key, fileBytes)
	if err != nil {
		vc.logger.Error("Cannot upload to S3", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload the file"})
		return
	}

	url := vc.cdn + "/lesson/" + key
	err = vc.repo.InsertVideoInfo(handsign, url)
	if err != nil {
		vc.logger.Error("Cannot insert into video table: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert to the table"})
	}

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
