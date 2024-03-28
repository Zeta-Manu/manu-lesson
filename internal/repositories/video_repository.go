package repositories

import (
	"github.com/Zeta-Manu/manu-lesson/internal/adapters/db"
	"github.com/Zeta-Manu/manu-lesson/internal/adapters/s3"
	"github.com/Zeta-Manu/manu-lesson/internal/domain"
)

var _ IVideoRepository = &VideoRepository{}

type IVideoRepository interface {
	GetVideo(id string) (*domain.Video, error)
	PostVideo(key string, file []byte) error
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

func (repo *VideoRepository) GetVideo(id string) (*domain.Video, error) {
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

func (repo *VideoRepository) PostVideo(key string, file []byte) error {
	err := repo.s3Adpter.PutObject(key, file)
	if err != nil {
		return err
	}
	return nil
}

func (repo *VideoRepository) GetAllVideo() ([]*domain.Video, error) {
	query := "SELECT id, handsign, url FROM video"
	rows, err := repo.dbAdapter.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []*domain.Video
	for rows.Next() {
		var video domain.Video
		err = rows.Scan(&video.ID, &video.HandSign, &video.VideoURL)
		if err != nil {
			return nil, err
		}
		videos = append(videos, &video)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}
