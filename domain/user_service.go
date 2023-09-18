package domain

import (
	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/internal/models"
	repository "github.com/mateusmlo/jornada-milhas/internal/repositories"
)

// IUserService interface for managing user resources
type IUserService interface {
	GetUserByUUID(id uuid.UUID) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByEmail(email string) (models.User, error)
	CreateUser(models.User) error
	UpdateUser(models.User) error
	DeactivateUser(id uuid.UUID) error
}

type UserService struct {
	repo repository.UserRepository
}

// NewUserService creates new userService
func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

// GetAllUsers returns all registered users
func (us *UserService) GetAllUsers() ([]*models.User, error) {
	users, err := us.repo.GetAllUsers()

	return users, err
}

// GetUserByUUID gets user by uuid PK
func (us *UserService) GetUserByUUID(id uuid.UUID) (*models.User, error) {
	user, err := us.repo.FindByUUID(id)

	return user, err
}

// CreateUser creates new user
func (us *UserService) CreateUser(u models.User) error {
	err := us.repo.CreateUser(u)

	return err
}
