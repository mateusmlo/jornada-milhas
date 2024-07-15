package service

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/internal/models"
	"github.com/mateusmlo/jornada-milhas/utils"
	"golang.org/x/crypto/bcrypt"
)

// AuthPayload authentication credentials payload
type AuthPayload struct {
	AccessToken string `json:"access_token"`
}

type authService struct {
	us UserService
	tu utils.TokenUtils
}

// NewAuthService creates new auth service
func NewAuthService(us UserService, tu utils.TokenUtils) AuthService {
	return &authService{
		us: us,
		tu: tu,
	}
}

// CreateSession attempts to login user and returns JWT pair
func (as *authService) CreateSession(payload dto.AuthDTO, user *models.User) (*AuthPayload, error) {
	validPassword := as.validatePasswordHash(payload.Password, user.Password)
	if !validPassword {
		return nil, errors.New("invalid email or password")
	}

	aTkn, err := as.GenerateTokenPair(user.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &AuthPayload{
		AccessToken: aTkn,
	}, nil
}

func (as *authService) validatePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func (as *authService) GenerateTokenPair(userID uuid.UUID) (string, error) {
	tkn, err := as.tu.GenerateAccessToken(userID)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tkn, nil
}
