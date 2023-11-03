package service

import (
	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/internal/models"
	repository "github.com/mateusmlo/jornada-milhas/internal/repositories"
)

type ReviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(r repository.ReviewRepository) *ReviewService {
	return &ReviewService{
		repo: r,
	}
}

func (rs *ReviewService) CreateReview(r *dto.NewReviewDTO, userID uuid.UUID) error {
	err := rs.repo.CreateReview(*r, userID)

	return err
}

func (rs *ReviewService) FindByUUID(id string, userID uuid.UUID) (*models.Review, error) {
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

func (rs *ReviewService) GetUserReviews(userID uuid.UUID) (*[]models.Review, error) {
	revs, err := rs.repo.GetUserReviews(userID)
	if err != nil {
		return nil, err
	}

	return revs, nil
}

func (rs *ReviewService) UpdateReview(id string, payload dto.UpdateReviewDTO, userID uuid.UUID) error {
	rID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	err = rs.repo.UpdateReview(payload, rID, userID)

	return err
}

func (rs *ReviewService) DeleteReview(reviewID string, userID uuid.UUID) (int64, error) {
	rID, err := uuid.Parse(reviewID)
	if err != nil {
		return 0, err
	}

	res, err := rs.repo.DeleteReview(userID, rID)

	return res, err
}
