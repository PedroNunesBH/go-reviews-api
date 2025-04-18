package entity 

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewReview(t *testing.T) {
	review, err := NewReview("Restaurante excelente", 4.8, "Bar do José")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "Restaurante excelente", review.Description)
	assert.Equal(t, 4.8, review.Rating)
	assert.Equal(t, "Bar do José", review.RestaurantID)
	assert.NotNil(t, review.CreatedAt)
}

func TestInvalidRating(t *testing.T) {
	_, err := NewReview("Restaurante excelente", 10.0, "Bar do José")
	assert.Equal(t, ErrInvalidRating, err)
	_, err = NewReview("Restaurante muito bom", -8.20, "Bar da Claúdia")
	assert.Equal(t, ErrInvalidRating, err)
}

func TestInvalidDescription(t *testing.T) {
	desc := "A comida estava excelente e o atendimento foi ainda melhor. " +
		"Com certeza voltarei mais vezes!"
	
	_, err := NewReview(desc, 4.98, "Bar do José")
	assert.Equal(t, ErrInvalidDescription, err)
}