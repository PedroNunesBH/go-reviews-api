package entity 

import (
	"go-reviews-api/pkg/entity"
	"errors"
)

var ErrInvalidCnpj = errors.New("invalid CNPJ: must have 14 characters")

type Restaurant struct {
	ID		entity.ID	`json:"id"`
	Name 	string	`json:"name"`
	Cnpj 	string	`json:"cnpj"`
	Address string	`json:"address"`
}

func NewRestaurant(name, cnpj, address string) (*Restaurant, error) {
	restaurant := &Restaurant{
		ID: entity.NewID(),
		Name: name,
		Cnpj: cnpj,
		Address: address,
	}
	err := restaurant.ValidateRestaurant()
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}

func (r *Restaurant) ValidateRestaurant() error {
	if len(r.Cnpj) < 14 {
		return ErrInvalidCnpj
	}
	return nil
}