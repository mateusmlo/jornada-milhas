package models

// Testimonial db model
type Testimonial struct {
	BaseModel
	Photo  string `gorm:"size:255;not null" json:"photo" validate:"required"`
	User   User   `gorm:"not null;foreignKey:UserID" json:"user" validate:"required"`
	UserID uint   `gorm:"column:user_id" json:"-"`
}
