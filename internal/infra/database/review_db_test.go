package database

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	"gorm.io/gorm"
	"github.com/glebarez/sqlite"
	"errors"
)

func TestCreateReview(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.AutoMigrate(&entity.Review{})
	if err != nil {
		t.Error()
	}
	review, err := entity.NewReview("Restaraunte excelente", 4.32, "jeoqu0p-19a-189361n")
	assert.Nil(t, err)
	reviewDB := NewReviewDB(db)
	err = reviewDB.CreateReview(review)
	assert.Nil(t, err)

	var reviewFound entity.Review
	err = db.First(&reviewFound, "id = ?", review.ID).Error
	assert.Nil(t, err)

	assert.Equal(t, review.Description, reviewFound.Description)
	assert.Equal(t, review.Rating, reviewFound.Rating)
	assert.Equal(t, review.RestaurantID, reviewFound.RestaurantID)
}

func TestFindAllReviews(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}
	db.AutoMigrate(&entity.Review{})
	reviewDB := NewReviewDB(db)
	review, err := entity.NewReview("Restaraunte excelente", 4.32, "jeoqu0p-19a-189361n")
	assert.Nil(t, err)
	err = reviewDB.CreateReview(review)
	assert.Nil(t, err)

	reviews, err := reviewDB.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(reviews))
}

func TestFindReviewByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}
	db.AutoMigrate(&entity.Review{})
	reviewDB := NewReviewDB(db)
	review, err := entity.NewReview("Restaraunte excelente", 4.32, "jeoqu0p-19a-189361n")
	assert.Nil(t, err)
	err = reviewDB.CreateReview(review)
	assert.Nil(t, err)
	reviewFound, err := reviewDB.FindReviewByID(review.ID)
	assert.Nil(t, err)

	assert.Equal(t, review.ID, reviewFound.ID)
	assert.Equal(t, review.Description, reviewFound.Description)
	assert.Equal(t, review.Rating, reviewFound.Rating)
}

func TestDeleteReview(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}
	db.AutoMigrate(&entity.Review{})
	review, err := entity.NewReview("Restaraunte excelente", 4.32, "jeoqu0p-19a-189361n")
	assert.Nil(t, err)
	reviewDB := NewReviewDB(db)
	err = reviewDB.CreateReview(review)
	assert.Nil(t, err)

	err = reviewDB.DeleteReview(review.ID)
	assert.Nil(t, err)
	_, err = reviewDB.FindReviewByID(review.ID)

	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestUpdateReview(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}
	db.AutoMigrate(&entity.Review{})
	review, err := entity.NewReview("Restaraunte excelente", 4.32, "jeoqu0p-19a-189361n")
	assert.Nil(t, err)
	reviewDB := NewReviewDB(db)
	err = reviewDB.CreateReview(review)
	assert.Nil(t, err)

	review.Description = "Restaurante agradável"
	review.Rating = 3.6
	err = reviewDB.UpdateReview(review)
	assert.Nil(t, err)
	reviewFound, err := reviewDB.FindReviewByID(review.ID)
	assert.Nil(t, err)

	assert.Equal(t, review.ID, reviewFound.ID)
	assert.Equal(t, "Restaurante agradável", reviewFound.Description)
	assert.Equal(t, 3.6, reviewFound.Rating)
}