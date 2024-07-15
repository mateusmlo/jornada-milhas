package service

import (
	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/internal/models"
)

// UserService provides user resources
type UserService interface {
	GetAllUsers() ([]*models.User, error)
	GetUserByUUID(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	CreateUser(u *dto.NewUserDTO) (*models.User, error)
	UpdateUser(id string, payload dto.UpdateUserDTO) error
	DeactivateUser(id string) (int64, error)
}

// ReviewService provides review resources
type ReviewService interface {
	CreateReview(r *dto.NewReviewDTO, userID uuid.UUID) error
	FindByUUID(id string, userID uuid.UUID) (*models.Review, error)
	GetUserReviews(userID uuid.UUID) (*[]models.Review, error)
	UpdateReview(id string, payload dto.UpdateReviewDTO, userID uuid.UUID) error
	DeleteReview(reviewID string, userID uuid.UUID) (int64, error)
}

// AuthService provides authentication resources
type AuthService interface {
	CreateSession(payload dto.AuthDTO, user *models.User) (*AuthPayload, error)
	validatePasswordHash(password, hash string) bool
	GenerateTokenPair(userID uuid.UUID) (string, error)
}
