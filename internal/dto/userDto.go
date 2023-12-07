package dto

// AuthDTO basic authentication data
type AuthDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// NewUserDTO struct for creating a new user
type NewUserDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=70"`
}

// UpdateUserDTO fields a user is allowed to update
type UpdateUserDTO struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty" validate:"email"`
	Password string `json:"password,omitempty" validate:"min=8,max=70"`
}
