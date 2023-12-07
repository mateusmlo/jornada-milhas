package repository

import (
	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/internal/models"
)

type UserRepository interface {
	GetAllUsers() ([]*models.User, error)
	FindByUUID(id uuid.UUID) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	CreateUser(u *dto.NewUserDTO) error
	UpdateUser(id uuid.UUID, u dto.UpdateUserDTO) error
	DeactivateUser(id uuid.UUID) (int64, error)
}

type ReviewRepository interface {
	CreateReview(r dto.NewReviewDTO, userID uuid.UUID) error
	FindByUUID(id, userID uuid.UUID) (*models.Review, error)
	UpdateReview(r dto.UpdateReviewDTO, id, userID uuid.UUID) error
	DeleteReview(userID, id uuid.UUID) (int64, error)
	GetUserReviews(userID uuid.UUID) (*[]models.Review, error)
}
