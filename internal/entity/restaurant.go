package entity 

import (
	"github.com/PedroNunesBH/go-reviews-api/pkg/entity"
)

type Restaraunt struct {
	ID		entity.ID	`json:"id"`
	Name 	string	`json:"name"`
	Cnpj 	string	`json:"cnpj"`
	Address string	`json:"address"`
}

func NewRestaraunt(name, cnpj, address string) (*Restaraunt, error) {
	return &Restaraunt{
		ID: entity.NewID(),
		Name: name,
		Cnpj: cnpj,
		Address: address,
	}, nil
}