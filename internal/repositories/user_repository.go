package repository

import (
	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/config"
	"github.com/mateusmlo/jornada-milhas/internal/models"
	"github.com/mateusmlo/jornada-milhas/tools"
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

	res := ur.Database.First(&user, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

// CreateUser creates new user
func (ur *UserRepository) CreateUser(u models.User) error {
	user := models.User{
		Name:  u.Name,
		Email: u.Email,
	}

	hashPassword, err := tools.HashPassword(u.Password)
	if err != nil {
		return err
	}

	user.Password = hashPassword

	tx := ur.Database.Begin()

	defer func() {
		RecoverPanic(ur.logger)
		tx.Rollback()
	}()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
