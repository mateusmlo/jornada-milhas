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
	Name     string    `gorm:"not null,size:100" json:"name"`
	Email    string    `gorm:"not null,size 100" json:"email"`
	Password string    `gorm:"size:70,not null" json:"-"`
	Reviews  []*Review `json:"-"`
}

func (u *User) BeforeSave(_ *gorm.DB) (err error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
		return
	}

	u.Password = string(hashPassword)
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return
}
