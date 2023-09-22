package models

import (
	"fmt"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User db model
type User struct {
	BaseModel
	Name     string    `gorm:"not null,size:100" json:"name" validate:"required"`
	Email    string    `gorm:"not null,size 100" json:"email" validate:"required"`
	Password string    `gorm:"size:255,not null" json:"-" validate:"required"`
	Reviews  []*Review `gorm:"foreignKey:ID" json:"reviews"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
		return
	}

	u.Password = string(hashPassword)
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return
}
