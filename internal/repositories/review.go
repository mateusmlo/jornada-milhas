package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/internal/models"
	"gorm.io/gorm"
)

// ReviewRepository struct
type reviewRepository struct {
	DB *gorm.DB
}

// NewReviewRepository new repo instance
func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{
		DB: db,
	}
}

// CreateReview creates new review
func (rr *reviewRepository) CreateReview(r dto.NewReviewDTO, userID uuid.UUID) error {
	review := models.Review{
		Review: r.Review,
		Photo:  r.Photo,
		UserID: userID,
	}

	tx := rr.DB.Begin()

	defer func() {
		RecoverPanic(tx.Statement.Context)
		tx.Rollback()
	}()

	if err := tx.Create(&review).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	return nil
}

func (rr *reviewRepository) FindByUUID(id, userID uuid.UUID) (*models.Review, error) {
	defer RecoverPanic(rr.DB.Statement.Context)

	var review models.Review

	if err := rr.DB.Preload("User").Where("user_id = ? AND id = ?", userID, id).First(&review).Error; err != nil {
		return nil, err
	}

	return &review, nil
}

func (rr *reviewRepository) UpdateReview(r dto.UpdateReviewDTO, id, userID uuid.UUID) error {
	review, err := rr.FindByUUID(id, userID)
	if err != nil {
		return err
	}

	tx := rr.DB.Begin()

	defer func() {
		RecoverPanic(tx.Statement.Context)
		tx.Rollback()
	}()

	if err := tx.Where(&review).Assign(&r).FirstOrCreate(&review).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	return nil
}

func (rr *reviewRepository) DeleteReview(userID, id uuid.UUID) (int64, error) {
	r, err := rr.FindByUUID(id, userID)
	if err != nil {
		return 0, err
	}

	res := rr.DB.Delete(&r)
	if res.Error != nil {
		return 0, err
	}

	return res.RowsAffected, nil
}

func (rr *reviewRepository) GetUserReviews(userID uuid.UUID) (*[]models.Review, error) {
	defer RecoverPanic(rr.DB.Statement.Context)

	var reviews []models.Review

	if err := rr.DB.Where("user_id = ?", userID).Find(&reviews).Error; err != nil {
		return nil, err
	}

	return &reviews, nil
}
