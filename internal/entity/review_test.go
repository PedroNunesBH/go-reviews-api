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