// Package dto contains structs to parse resource creation and update
package dto

import "github.com/google/uuid"

// NewReviewDTO contains data needed for a new review
type NewReviewDTO struct {
	Review string `json:"review" binding:"required,max=255"`
	Photo  string `json:"photo" binding:"required,max=255"`
}

// UpdateReviewDTO contains data needed to update review
type UpdateReviewDTO struct {
	ID     uuid.UUID `json:"id" binding:"required"`
	Review string    `json:"review,omitempty" binding:"max=255"`
	Photo  string    `json:"photo,omitempty" binding:"max=255"`
}
