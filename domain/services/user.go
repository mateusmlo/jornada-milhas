package service

import (
	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/internal/models"
	repository "github.com/mateusmlo/jornada-milhas/internal/repositories"
)

type userService struct {
	repo repository.UserRepository
}

// NewUserService creates new userService
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// GetAllUsers returns all registered users
func (us *userService) GetAllUsers() ([]*models.User, error) {
	users, err := us.repo.GetAllUsers()

	return users, err
}

// GetUserByUUID gets user by uuid PK
func (us *userService) GetUserByUUID(id string) (*models.User, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	user, err := us.repo.FindByUUID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindByEmail finds user by email
func (us *userService) FindByEmail(email string) (*models.User, error) {
	user, err := us.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser creates new user
func (us *userService) CreateUser(u *dto.NewUserDTO) error {
	err := us.repo.CreateUser(u)

	return err
}

// UpdateUser updates a user
func (us *userService) UpdateUser(id string, payload dto.UpdateUserDTO) error {
	userID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	err = us.repo.UpdateUser(userID, payload)

	return err
}

// DeactivateUser deactivates a user - it does NOT delete!
func (us *userService) DeactivateUser(id string) (int64, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return 0, err
	}

	res, err := us.repo.DeactivateUser(userID)

	return res, err
}
