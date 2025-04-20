package entity 

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/PedroNunesBH/go-reviews-api/pkg/entity"
)

func TestCreateNewReview(t *testing.T) {
	restaurantID := entity.NewID()
	review, err := NewReview("Restaurante excelente", 4.8, restaurantID)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "Restaurante excelente", review.Description)
	assert.Equal(t, 4.8, review.Rating)
	assert.Equal(t, restaurantID, review.RestaurantID)
	assert.NotNil(t, review.CreatedAt)
}

func TestInvalidRating(t *testing.T) {
	_, err := NewReview("Restaurante excelente", 10.0, entity.NewID())
	assert.Equal(t, ErrInvalidRating, err)
	_, err = NewReview("Restaurante muito bom", -8.20, entity.NewID())
	assert.Equal(t, ErrInvalidRating, err)
}

func TestInvalidDescription(t *testing.T) {
	desc := "A comida estava excelente e o atendimento foi ainda melhor. " +
		"Com certeza voltarei mais vezes!"
	
	_, err := NewReview(desc, 4.98, entity.NewID())
	assert.Equal(t, ErrInvalidDescription, err)
}