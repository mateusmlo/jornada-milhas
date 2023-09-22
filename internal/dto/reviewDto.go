package dto

import "github.com/google/uuid"

type NewReviewDTO struct {
	Review string    `json:"review" binding:"required,max=255"`
	Photo  string    `json:"photo" binding:"required,max=255"`
	UserID uuid.UUID `json:"user" binding:"required,uuid"`
}

type UpdateReviewDTO struct {
	Review string    `json:"review,omitempty" binding:"required,max=255"`
	Photo  string    `json:"photo,omitempty" binding:"required,max=255"`
	UserID uuid.UUID `json:"user,omitempty" binding:"required,uuid"`
}
