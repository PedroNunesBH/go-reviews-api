package dto

import (
	"github.com/PedroNunesBH/go-reviews-api/pkg/entity"
)

type UserResponseDTO struct {
	ID  entity.ID    `json:"id"`
	Username  string `json:"username"`
	Email  string    `json:"email"`
}

type UserRequestDTO struct {
	Username string		`json:"username"`
	Email string		`json:"email"`
	Password string		`json:"password"`
}