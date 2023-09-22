package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/config"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/internal/models"
	"gorm.io/gorm"
)

// UserRepository DB structure
type UserRepository struct {
	DB     *gorm.DB
	logger *config.GormLogger
}

func RecoverPanic(ctx context.Context, logger *config.GormLogger) {
	if r := recover(); r != nil {
		logger.Warn(ctx, "😵 Recovered from panic!")
	}
}

// NewUserRepository creates a new user repository
func NewUserRepository(logger *config.GormLogger, db *gorm.DB) UserRepository {
	return UserRepository{
		DB:     db,
		logger: logger,
	}
}

// GetAllUsers get all registered users
func (ur *UserRepository) GetAllUsers() ([]*models.User, error) {
	defer RecoverPanic(ur.DB.Statement.Context, ur.logger)

	var users []*models.User

	res := ur.DB.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

// FindByUUID finds user by PK
func (ur *UserRepository) FindByUUID(id uuid.UUID) (*models.User, error) {
	defer RecoverPanic(ur.DB.Statement.Context, ur.logger)

	var user models.User

	if err := ur.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByEmail finds user by email
func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	defer RecoverPanic(ur.DB.Statement.Context, ur.logger)

	var user models.User

	if err := ur.DB.First(&user, models.User{Email: email}).Error; err != nil {
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

	tx := ur.DB.Begin()

	defer func() {
		RecoverPanic(tx.Statement.Context, ur.logger)
		tx.Rollback()
	}()

	if err := tx.Create(&user).Error; err != nil {
		ur.logger.Error(tx.Statement.Context, err.Error())
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		ur.logger.Error(tx.Statement.Context, err.Error())
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

	tx := ur.DB.Begin()

	defer func() {
		RecoverPanic(tx.Statement.Context, ur.logger)
		tx.Rollback()
	}()

	if err := tx.Where(&user).Assign(&u).FirstOrCreate(&user).Error; err != nil {
		ur.logger.Error(tx.Statement.Context, err.Error())
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		ur.logger.Error(tx.Statement.Context, err.Error())
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

	res := ur.DB.Delete(&user)
	if res.Error != nil {
		return 0, err
	}

	return res.RowsAffected, nil
}
