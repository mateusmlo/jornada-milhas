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

func (rs *ReviewService) CreateReview(r *dto.NewReviewDTO) error {
	err := rs.repo.CreateReview(*r)

	return err
}

func (rs *ReviewService) FindByUUID(id string) (*models.Review, error) {
	reviewID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	review, err := rs.repo.FindByUUID(reviewID)
	if err != nil {
		return nil, err
	}

	return review, nil
}

func (rs *ReviewService) UpdateReview(id string, payload dto.UpdateReviewDTO) error {
	reviewID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	err = rs.repo.UpdateReview(reviewID, payload)

	return err
}
