package service

import (
	"errors"
	"fmt"

	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/internal/models"
	"github.com/mateusmlo/jornada-milhas/tools"
	"golang.org/x/crypto/bcrypt"
)

// AuthPayload authentication credentials payload
type AuthPayload struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type authService struct {
	us UserService
	rs RefreshService
	tu tools.TokenUtils
}

// NewAuthService creates new auth service
func NewAuthService(us UserService, rs RefreshService, tu tools.TokenUtils) AuthService {
	return &authService{
		us: us,
		rs: rs,
		tu: tu,
	}
}

// CreateSession attempts to login user and returns JWT pair
func (as *authService) CreateSession(payload dto.AuthDTO, user *models.User) (*AuthPayload, error) {
	validPassword := as.validatePasswordHash(payload.Password, user.Password)
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

// Logout attempts to logout user
func (as *authService) Logout(user *models.User) bool {
	return as.rs.DeleteRefreshToken(user.ID.String())
}

func (as *authService) validatePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
