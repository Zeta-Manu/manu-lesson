package repositories

import (
	"github.com/Zeta-Manu/manu-lesson/internal/adapters/db"
	"github.com/Zeta-Manu/manu-lesson/internal/adapters/s3"
	"github.com/Zeta-Manu/manu-lesson/internal/domain"
)

var _ IVideoRepository = &VideoRepository{}

type IVideoRepository interface {
	GetVideo(id int) (*domain.Video, error)
	PostVideo(file byte) error
	GetAllVideo() ([]*domain.Video, error)
}

type VideoRepository struct {
	dbAdapter *db.Database
	s3Adpter  *s3.S3Adapter
}

func NewVideoRepository(dbAdapter *db.Database, S3Adapter *s3.S3Adapter) *VideoRepository {
	return &VideoRepository{
		dbAdapter: dbAdapter,
		s3Adpter:  S3Adapter,
	}
}

// NOTE: Get Video from video table
func (repo *VideoRepository) GetVideo(id int) (*domain.Video, error) {
	query := "SELECT id, handsign, url FROM video WHERE id = ?"
	rows, err := repo.dbAdapter.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var video domain.Video
	if rows.Next() {
		err = rows.Scan(&video.ID, &video.HandSign, &video.VideoURL)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	return &video, nil
}

// NOTE: Upload a file to s3 and return CloudFront URL
func (repo *VideoRepository) PostVideo(file byte) error {
	return nil
}

func (repo *VideoRepository) GetAllVideo() ([]*domain.Video, error) {
	return nil, nil
}
