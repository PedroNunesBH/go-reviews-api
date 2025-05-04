package dto

import (
	"github.com/PedroNunesBH/go-reviews-api/pkg/entity"
)

type UserResponsDTO struct {
	ID  entity.ID    `json:"id"`
	Username  string `json:"username"`
	Email  string    `json:"email"`
}