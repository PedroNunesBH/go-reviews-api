package entity

import (
	"time"
	"github.com/PedroNunesBH/go-reviews-api/pkg/entity"
)

type Review struct {
	ID           entity.ID  `json:"id"`
	Description  string		`json:"description"`
	Rating       float64	`json:"rating"`
	CreatedAt 	 time.Time	`json:"created_at"`	
	RestaurantID string		`json:"restaurant_id"`
}

func NewReview(description string, rating float64, restaurantID string) (*Review, error) {
	return &Review {
		ID: entity.NewID(),
		Description: description,
		Rating: rating,
		CreatedAt: time.Now(),
		RestaurantID: restaurantID,
	}, nil
}