package models

// User db model
type User struct {
	BaseModel
	Name         string         `gorm:"not null;size:100" json:"name" validate:"required"`
	Email        string         `gorm:"not null;size 100" json:"email" validate:"required"`
	Password     string         `gorm:"size:30;not null" json:"-" validate:"required"`
	Testimonials []*Testimonial `gorm:"foreignKey:ID" json:"testimonials"`
}
