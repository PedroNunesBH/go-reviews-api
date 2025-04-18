package entity 

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewRestaurant(t *testing.T) {
	restaurant, err := NewRestaraunt("Bar do José", "192415627181", "Rua das Flores - Belo Horizonte - MG")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "Bar do José", restaurant.Name)
	assert.Equal(t, "192415627181", restaurant.Cnpj)
	assert.Equal(t, "Rua das Flores - Belo Horizonte - MG", restaurant.Address)
}