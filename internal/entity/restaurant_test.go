package entity 

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewRestaurant(t *testing.T) {
	restaurant, err := NewRestaurant("Bar do José", "19241562718190", "Rua das Flores - Belo Horizonte - MG")
	assert.Nil(t, err)
	assert.Equal(t, "Bar do José", restaurant.Name)
	assert.Equal(t, "19241562718190", restaurant.Cnpj)
	assert.Equal(t, "Rua das Flores - Belo Horizonte - MG", restaurant.Address)
}

func TestValidateCnpjWithInvalid(t *testing.T) {
	_, err := NewRestaurant("Bar do José", "192415627181", "Rua das Flores - Belo Horizonte - MG")
	assert.Equal(t, ErrInvalidCnpj, err)
}
