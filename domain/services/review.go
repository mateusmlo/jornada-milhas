package service

import (
	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/internal/models"
	repo "github.com/mateusmlo/jornada-milhas/internal/repositories"
)

type reviewService struct {
	repo repo.ReviewRepository
}

// NewReviewService returns new service instance
func NewReviewService(r repo.ReviewRepository) ReviewService {
	return &reviewService{
		repo: r,
	}
}

func (rs *reviewService) CreateReview(r *dto.NewReviewDTO, userID uuid.UUID) error {
	err := rs.repo.CreateReview(*r, userID)

	return err
}

func (rs *reviewService) FindByUUID(id string, userID uuid.UUID) (*models.Review, error) {
	reviewID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	review, err := rs.repo.FindByUUID(reviewID, userID)
	if err != nil {
		return nil, err
	}

	return review, nil
}

func (rs *reviewService) GetUserReviews(userID uuid.UUID) (*[]models.Review, error) {
	revs, err := rs.repo.GetUserReviews(userID)
	if err != nil {
		return nil, err
	}

	return revs, nil
}

func (rs *reviewService) UpdateReview(id string, payload dto.UpdateReviewDTO, userID uuid.UUID) error {
	rID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	err = rs.repo.UpdateReview(payload, rID, userID)

	return err
}

func (rs *reviewService) DeleteReview(reviewID string, userID uuid.UUID) (int64, error) {
	rID, err := uuid.Parse(reviewID)
	if err != nil {
		return 0, err
	}

	res, err := rs.repo.DeleteReview(userID, rID)

	return res, err
}
