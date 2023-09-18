package domain

import "github.com/mateusmlo/jornada-milhas/internal/models"

// AuthService authorization functions
type AuthService interface {
	Login(tkn string) (bool, error)
	GenerateJWT(models.User) string
}
