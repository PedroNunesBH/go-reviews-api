package dto

import (
	"go-reviews-api/pkg/entity"
)

type ReviewDTO struct {
	Description   string	`json:"description"`
	Rating        float64	`json:"rating"`
	RestaurantID  entity.ID	`json:"restaurant_id"`
}