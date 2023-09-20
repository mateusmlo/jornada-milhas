package models

// User db model
type User struct {
	BaseModel
	Name         string         `gorm:"not null,size:100" json:"name" validate:"required"`
	Email        string         `gorm:"not null,size 100" json:"email" validate:"required"`
	Password     string         `gorm:"size:255,not null" json:"-" validate:"required"`
	Testimonials []*Testimonial `gorm:"foreignKey:ID" json:"testimonials"`
}

// UpdateUserDTO fields a user is allowed to update
type UpdateUserDTO struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
