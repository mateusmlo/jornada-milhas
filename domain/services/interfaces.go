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
	CreateUser(u *dto.NewUserDTO) error
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

// RefreshService provides refresh token resources
type RefreshService interface {
	GetRefreshToken(userID string) (string, error)
	SetRefreshToken(tkn, userID string) error
	DeleteRefreshToken(userID string) bool
}

// AuthService provides authentication resources
type AuthService interface {
	CreateSession(payload dto.AuthDTO, user *models.User) (*AuthPayload, error)
	Logout(user *models.User) bool
	validatePasswordHash(password, hash string) bool
}
