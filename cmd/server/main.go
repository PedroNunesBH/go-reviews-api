package main

import (
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/PedroNunesBH/go-reviews-api/internal/webserver/handlers"
	"github.com/PedroNunesBH/go-reviews-api/internal/infra/database"
)

func main() {
	db, err := gorm.Open(sqlite.Open("reviews.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Restaurant{}, &entity.Review{})

	restaurantRepo := database.NewRestaurantDB(db)
	restaurantHandler := handlers.NewRestaurantHandler(restaurantRepo)

	r := chi.NewRouter()
	r.Route("/restaurants", func (r chi.Router) {
		r.Post("/", restaurantHandler.CreateRestaurant)
		r.Get("/{id}", restaurantHandler.GetRestaurant)
		r.Delete("/{id}", restaurantHandler.DeleteRestaurant)
	})
	http.ListenAndServe(":8000", r)

}