package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/config"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/internal/models"
	"gorm.io/gorm"
)

// UserRepository database structure
type UserRepository struct {
	Database *gorm.DB
	logger   config.Logger
}

func RecoverPanic(logger config.Logger) {
	if r := recover(); r != nil {
		logger.Warn("😵 Recovered from panic!")
	}
}

// NewUserRepository creates a new user repository
func NewUserRepository(logger config.Logger, db *gorm.DB) UserRepository {
	return UserRepository{
		Database: db,
		logger:   logger,
	}
}

// GetAllUsers get all registered users
func (ur *UserRepository) GetAllUsers() ([]*models.User, error) {
	defer RecoverPanic(ur.logger)

	var users []*models.User

	res := ur.Database.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

// FindByUUID finds user by PK
func (ur *UserRepository) FindByUUID(id uuid.UUID) (*models.User, error) {
	defer RecoverPanic(ur.logger)

	var user models.User

	if err := ur.Database.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByEmail finds user by email
func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	defer RecoverPanic(ur.logger)

	var user models.User

	if err := ur.Database.First(&user, models.User{Email: email}).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser creates new user
func (ur *UserRepository) CreateUser(u dto.NewUserDTO) error {
	user := models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}

	tx := ur.Database.Begin()

	defer func() {
		RecoverPanic(ur.logger)
		tx.Rollback()
	}()

	if err := tx.Create(&user).Error; err != nil {
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

// UpdateUser updates user data
func (ur *UserRepository) UpdateUser(id uuid.UUID, u dto.UpdateUserDTO) error {
	user, err := ur.FindByUUID(id)
	if err != nil {
		return err
	}

	tx := ur.Database.Begin()

	defer func() {
		RecoverPanic(ur.logger)
		tx.Rollback()
	}()

	if err := tx.Where(&user).Assign(&u).FirstOrCreate(&user).Error; err != nil {
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

// DeactivateUser deactivates a user
func (ur *UserRepository) DeactivateUser(id uuid.UUID) (int64, error) {
	user, err := ur.FindByUUID(id)
	if err != nil {
		return 0, err
	}

	res := ur.Database.Delete(&user)
	if res.Error != nil {
		return 0, err
	}

	return res.RowsAffected, nil
}
