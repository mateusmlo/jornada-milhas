package dto

import "github.com/google/uuid"

type NewReviewDTO struct {
	Review string `json:"review" binding:"required,max=255"`
	Photo  string `json:"photo" binding:"required,max=255"`
	UserID uuid.UUID
}

type UpdateReviewDTO struct {
	Review string `json:"review,omitempty" binding:"-,max=255"`
	Photo  string `json:"photo,omitempty" binding:"-,max=255"`
	UserID uuid.UUID
}
