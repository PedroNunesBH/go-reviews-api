package main

import (
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("reviews.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Restaurant{}, &entity.Review{})
}