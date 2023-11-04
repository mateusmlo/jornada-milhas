package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/internal/models"
	"gorm.io/gorm"
)

// UserRepository DB structure
type UserRepository struct {
	DB *gorm.DB
}

func RecoverPanic(ctx context.Context) {
	if r := recover(); r != nil {
		fmt.Println(ctx, "😵 Recovered from panic!")
	}
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		DB: db,
	}
}

// GetAllUsers get all registered users
func (ur *UserRepository) GetAllUsers() ([]*models.User, error) {
	defer RecoverPanic(ur.DB.Statement.Context)

	var users []*models.User

	res := ur.DB.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

// FindByUUID finds user by PK
func (ur *UserRepository) FindByUUID(id uuid.UUID) (*models.User, error) {
	defer RecoverPanic(ur.DB.Statement.Context)

	var user models.User

	if err := ur.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByEmail finds user by email
func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	defer RecoverPanic(ur.DB.Statement.Context)

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
		RecoverPanic(tx.Statement.Context)
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

	tx := ur.DB.Begin()

	defer func() {
		RecoverPanic(tx.Statement.Context)
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

	res := ur.DB.Delete(&user)
	if res.Error != nil {
		return 0, err
	}

	return res.RowsAffected, nil
}
