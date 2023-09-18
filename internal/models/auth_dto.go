package models

// AuthDTO basic authentication data
type AuthDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
