package service

import (
	"errors"
	"fmt"

	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/internal/models"
	"github.com/mateusmlo/jornada-milhas/tools"
	"golang.org/x/crypto/bcrypt"
)

type AuthPayload struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// AuthService provides authentication resources
type AuthService struct {
	us *UserService
	rs *RefreshService
	tu *tools.TokenUtils
}

// NewAuthService creates new auth service
func NewAuthService(us *UserService, rs *RefreshService, tu *tools.TokenUtils) *AuthService {
	return &AuthService{
		us: us,
		rs: rs,
		tu: tu,
	}
}

func (as *AuthService) CreateSession(payload dto.AuthDTO, user *models.User) (*AuthPayload, error) {
	validPassword := validatePasswordHash(payload.Password, user.Password)
	if !validPassword {
		return nil, errors.New("Invalid email or password")
	}

	tkn, err := as.tu.GenerateAccessToken(user.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rTkn, err := as.tu.GenerateRefreshToken(user.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = as.rs.SetRefreshToken(rTkn, user.ID.String())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &AuthPayload{
		AccessToken:  tkn,
		RefreshToken: rTkn,
	}, nil
}

func (as *AuthService) Logout(user *models.User) bool {
	return as.rs.DeleteRefreshToken(user.ID.String())
}

func validatePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
