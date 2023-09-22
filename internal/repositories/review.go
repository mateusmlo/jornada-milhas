package repository

import (
	"github.com/mateusmlo/jornada-milhas/config"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/internal/models"
	"gorm.io/gorm"
)

// ReviewRepository struct
type ReviewRepository struct {
	DB     *gorm.DB
	logger *config.GormLogger
}

// NewReviewRepository new repo instance
func NewReviewRepository(logger *config.GormLogger, db *gorm.DB) ReviewRepository {
	return ReviewRepository{
		logger: logger,
		DB:     db,
	}
}

// CreateReview creates new review
func (rr *ReviewRepository) CreateReview(r dto.NewReviewDTO) error {
	review := models.Review{
		Review: r.Review,
		Photo:  r.Photo,
		UserID: r.UserID,
	}

	tx := rr.DB.Begin()

	defer func() {
		RecoverPanic(tx.Statement.Context, rr.logger)
		tx.Rollback()
	}()

	if err := tx.Create(&review).Error; err != nil {
		rr.logger.Error(tx.Statement.Context, err.Error())
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		rr.logger.Error(tx.Statement.Context, err.Error())
		tx.Rollback()
		return err
	}

	return nil
}
