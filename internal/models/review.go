package models

import "github.com/google/uuid"

// Review db model
type Review struct {
	BaseModel
	Review string    `gorm:"size:255;not null" json:"review" validate:"required"`
	Photo  string    `gorm:"size:255;not null" json:"photo" validate:"required"`
	User   User      `gorm:"not null;foreignKey:UserID" json:"user" validate:"required"`
	UserID uuid.UUID `gorm:"column:user_id" json:"-"`
}
