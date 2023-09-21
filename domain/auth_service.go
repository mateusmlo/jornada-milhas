package domain

import (
	"errors"
	"fmt"

	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/internal/models"
	"github.com/mateusmlo/jornada-milhas/tools"
	"golang.org/x/crypto/bcrypt"
)

// IAuthService interface
type IAuthService interface {
	Login(tkn string) (bool, error)
	GenerateJWT(models.User) string
}

// AuthService provides authentication resources
type AuthService struct {
	us *UserService
}

// NewAuthService creates new auth service
func NewAuthService(us *UserService) *AuthService {
	return &AuthService{
		us: us,
	}
}

func (as *AuthService) CreateSession(payload dto.AuthDTO, user *models.User) (string, error) {
	validPassword := validatePasswordHash(payload.Password, user.Password)
	if !validPassword {
		return "", errors.New("Invalid email or password")
	}

	tkn, err := tools.GenerateJWT(user.ID)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tkn, nil
}

func validatePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
