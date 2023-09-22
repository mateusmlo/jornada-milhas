package dto

// AuthDTO basic authentication data
type AuthDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// NewUserDTO struct for creating a new user
type NewUserDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UpdateUserDTO fields a user is allowed to update
type UpdateUserDTO struct {
	Name     string `json:"name,omitempty" binding:"required"`
	Email    string `json:"email,omitempty" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required"`
}
