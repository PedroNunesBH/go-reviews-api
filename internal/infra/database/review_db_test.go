package database

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	"gorm.io/gorm"
	"github.com/glebarez/sqlite"
	"errors"
	pkgEntity "github.com/PedroNunesBH/go-reviews-api/pkg/entity"
)

func TestCreateReview(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.AutoMigrate(&entity.Review{}, &entity.Restaurant{})
	if err != nil {
		t.Error()
	}
	restaurant, err := entity.NewRestaurant("Bar da Maria", "19020154829102", "Rua das Orquideas 1809")
	assert.Nil(t, err)
	restaurantDB := NewRestaurantDB(db)
	err = restaurantDB.CreateRestaurant(restaurant)
	assert.Nil(t, err)
	review, err := entity.NewReview("Restaraunte excelente", 4.32, restaurant.ID)
	assert.Nil(t, err)
	reviewDB := NewReviewDB(db)
	assert.Nil(t, err)
	err = reviewDB.CreateReview(review)
	assert.Nil(t, err)

	var reviewFound entity.Review
	err = db.First(&reviewFound, "id = ?", review.ID).Error
	assert.Nil(t, err)

	assert.Equal(t, review.Description, reviewFound.Description)
	assert.Equal(t, review.Rating, reviewFound.Rating)
	assert.Equal(t, review.RestaurantID, reviewFound.RestaurantID)
}

func TestCreateReviewWithAnInexistentRestaurantID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Restaurant{}, &entity.Review{})

	reviewDB := NewReviewDB(db)
	restaurantID, err := pkgEntity.ParseID("0b6bd722-7ee9-415c-826c-5e165b6781fe")
	assert.Nil(t, err)
	review, err := entity.NewReview("Restaraunte excelente", 4.32, restaurantID)
	assert.Nil(t, err)
	err = reviewDB.CreateReview(review)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestFindAllReviews(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}
	db.AutoMigrate(&entity.Review{}, &entity.Restaurant{})
	reviewDB := NewReviewDB(db)
	restaurant, err := entity.NewRestaurant("Bar da Maria", "19020154829102", "Rua das Orquideas 1809")
	assert.Nil(t, err)
	restaurantDB := NewRestaurantDB(db)
	err = restaurantDB.CreateRestaurant(restaurant)
	assert.Nil(t, err)
	review, err := entity.NewReview("Restaraunte excelente", 4.32, restaurant.ID)
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
	db.AutoMigrate(&entity.Review{}, &entity.Restaurant{})
	reviewDB := NewReviewDB(db)
	restaurant, err := entity.NewRestaurant("Bar da Maria", "19020154829102", "Rua das Orquideas 1809")
	assert.Nil(t, err)
	restaurantDB := NewRestaurantDB(db)
	err = restaurantDB.CreateRestaurant(restaurant)
	assert.Nil(t, err)
	review, err := entity.NewReview("Restaraunte excelente", 4.32, restaurant.ID)
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
	db.AutoMigrate(&entity.Review{}, &entity.Restaurant{}) 
	restaurant, err := entity.NewRestaurant("Bar da Maria", "19020154829102", "Rua das Orquideas 1809")
	assert.Nil(t, err)
	restaurantDB := NewRestaurantDB(db)
	err = restaurantDB.CreateRestaurant(restaurant)
	assert.Nil(t, err)
	review, err := entity.NewReview("Restaraunte excelente", 4.32, restaurant.ID)
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
	db.AutoMigrate(&entity.Review{}, &entity.Restaurant{})
	restaurant, err := entity.NewRestaurant("Bar da Maria", "19020154829102", "Rua das Orquideas 1809")
	assert.Nil(t, err)
	restaurantDB := NewRestaurantDB(db)
	err = restaurantDB.CreateRestaurant(restaurant)
	assert.Nil(t, err)
	review, err := entity.NewReview("Restaraunte excelente", 4.32, restaurant.ID)
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
