package entity

import (
	"time"
	"github.com/PedroNunesBH/go-reviews-api/pkg/entity"
	"errors"
)

var ErrInvalidRating = errors.New("rating must be between 0 and 5")
var ErrInvalidDescription = errors.New("description can't have more than 50 characters")

type Review struct {
	ID           entity.ID  `json:"id"`
	Description  string		`json:"description"`
	Rating       float64	`json:"rating"`
	CreatedAt 	 time.Time	`json:"created_at"`	
	RestaurantID string		`json:"restaurant_id"`
}

func NewReview(description string, rating float64, restaurantID string) (*Review, error) {
	review := &Review {
		ID: entity.NewID(),
		Description: description,
		Rating: rating,
		CreatedAt: time.Now(),
		RestaurantID: restaurantID,
	}
	err := review.ValidateReview()
	if err != nil {
		return nil, err
	}
	return review, err
}

func (r *Review) ValidateReview() error {
	if r.Rating < 0.0 || r.Rating > 5.0 {
		return ErrInvalidRating
	}
	if len(r.Description) > 50 {
		return ErrInvalidDescription
	}
	return nil
}