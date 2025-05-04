package main

import (
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/PedroNunesBH/go-reviews-api/internal/webserver/handlers"
	"github.com/PedroNunesBH/go-reviews-api/internal/infra/database"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/PedroNunesBH/go-reviews-api/docs"
)

// @title           Reviews API
// @version         1.0
// @description     This is an API for restaurant review management
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8000
// @BasePath  /

func main() {
	db, err := gorm.Open(sqlite.Open("reviews.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Restaurant{}, &entity.Review{}, &entity.User{})

	restaurantRepo := database.NewRestaurantDB(db)
	restaurantHandler := handlers.NewRestaurantHandler(restaurantRepo)

	reviewRepo := database.NewReviewDB(db)
	reviewHandler := handlers.NewReviewHandler(reviewRepo)

	userRepo := database.NewUserDB(db)
	userHandler := handlers.NewUserHandler(userRepo)

	r := chi.NewRouter()

	r.Route("/users", func (r chi.Router) {
		r.Get("/", userHandler.GetAllUsers)
		r.Post("/", userHandler.CreateUser)
		r.Get("/{id}", userHandler.GetUserByID)
		r.Delete("/{id}", userHandler.DeleteUser)
		r.Put("/{id}", userHandler.UpdateUser)
	})

	r.Route("/restaurants", func (r chi.Router) {
		r.Post("/", restaurantHandler.CreateRestaurant)
		r.Get("/{id}", restaurantHandler.GetRestaurant)
		r.Delete("/{id}", restaurantHandler.DeleteRestaurant)
		r.Get("/", restaurantHandler.GetAllRestaurants)
		r.Put("/{id}", restaurantHandler.UpdateRestaurant)
	})
	
	r.Route("/reviews", func (r chi.Router) {
		r.Post("/", reviewHandler.CreateReview)
		r.Get("/{id}", reviewHandler.GetReviewByID)
		r.Delete("/{id}", reviewHandler.DeleteReview)
		r.Get("/", reviewHandler.GetAllReviews)
		r.Put("/{id}", reviewHandler.UpdateReview)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	http.ListenAndServe(":8000", r)

}