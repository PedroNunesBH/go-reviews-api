package database

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"github.com/glebarez/sqlite"
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	"errors"
)

func TestCreateRestaurant(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}
	db.AutoMigrate(&entity.Restaurant{})
	restaurantDB := NewRestaurantDB(db)

	restaurant, err := entity.NewRestaurant("Bar da Maria", "19020154829102", "Rua das Orquideas 1809")
	assert.Nil(t, err)
	err = restaurantDB.CreateRestaurant(restaurant)
	assert.Nil(t, err)
	restaurantFound, err := restaurantDB.FindRestaurantByID(restaurant.ID)
	assert.Nil(t, err)

	assert.Equal(t, restaurant.ID, restaurantFound.ID)
	assert.Equal(t, restaurant.Name, restaurantFound.Name)
	assert.Equal(t, restaurant.Cnpj, restaurantFound.Cnpj)
	assert.Equal(t, restaurant.Address, restaurantFound.Address)
}

func TestDeleteRestaurant(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}
	db.AutoMigrate(&entity.Restaurant{})
	restaurantDB := NewRestaurantDB(db)

	restaurant, err := entity.NewRestaurant("Bar da Maria", "19020154829102", "Rua das Orquideas 1809")
	assert.Nil(t, err)
	err = restaurantDB.CreateRestaurant(restaurant)
	assert.Nil(t, err)
	err = restaurantDB.DeleteRestaurant(restaurant.ID)
	assert.Nil(t, err)
	_, err = restaurantDB.FindRestaurantByID(restaurant.ID)

	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestFindRestaurantById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}
	db.AutoMigrate(&entity.Restaurant{})
	restaurantDB := NewRestaurantDB(db)

	restaurant, err := entity.NewRestaurant("Bar da Maria", "19020154829102", "Rua das Orquideas 1809")
	assert.Nil(t, err)
	err = restaurantDB.CreateRestaurant(restaurant)
	assert.Nil(t, err)
	restaurantFound, err := restaurantDB.FindRestaurantByID(restaurant.ID)
	assert.Nil(t, err)

	assert.Equal(t, restaurant.ID, restaurantFound.ID)
	assert.Equal(t, restaurant.Name, restaurantFound.Name)
	assert.Equal(t, restaurant.Cnpj, restaurantFound.Cnpj)
	assert.Equal(t, restaurant.Address, restaurantFound.Address)
}

func TestFindAllRestaurants(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}
	db.AutoMigrate(&entity.Restaurant{})
	restaurantDB := NewRestaurantDB(db)

	restaurant, err := entity.NewRestaurant("Bar da Maria", "19020154829102", "Rua das Orquideas 1809")
	assert.Nil(t, err)
	err = restaurantDB.CreateRestaurant(restaurant)
	assert.Nil(t, err)
	secondRestaurant, err := entity.NewRestaurant("Bar do José", "19020154829982", "Rua Afonso Belo 1809")
	assert.Nil(t, err)
	err = restaurantDB.CreateRestaurant(secondRestaurant)
	assert.Nil(t, err)
	restaurants, err := restaurantDB.FindAllRestaurants()
	assert.Nil(t, err)

	assert.Equal(t, 2, len(restaurants))
}

func TestUpdateRestaurant(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}
	db.AutoMigrate(&entity.Restaurant{})
	restaurantDB := NewRestaurantDB(db)

	restaurant, err := entity.NewRestaurant("Bar da Maria", "19020154829102", "Rua das Orquideas 1809")
	assert.Nil(t, err)
	err = restaurantDB.CreateRestaurant(restaurant)
	assert.Nil(t, err)

	restaurant.Name = "Bar da Esquina"
	restaurant.Address = "Rua 7 Número 2"
	err = restaurantDB.UpdateRestaurant(restaurant)
	assert.Nil(t, err)
	restaurantFound, err := restaurantDB.FindRestaurantByID(restaurant.ID)
	assert.Nil(t, err)

	assert.Equal(t, restaurant.ID, restaurantFound.ID)
	assert.Equal(t, "Bar da Esquina", restaurantFound.Name)
	assert.Equal(t, "Rua 7 Número 2", restaurantFound.Address)
}