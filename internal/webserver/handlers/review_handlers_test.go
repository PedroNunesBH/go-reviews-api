package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	"github.com/PedroNunesBH/go-reviews-api/internal/infra/database"
	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ReviewHandlersTestSuit struct {
	suite.Suite
	ReviewHandler *ReviewHandler
	Review        *entity.Review
	Restaurant    *entity.Restaurant
	Router        http.Handler
}

func (suite *ReviewHandlersTestSuit) SetupTest() {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	suite.Nil(err)

	db.AutoMigrate(&entity.Review{}, &entity.Restaurant{})

	restaurant, err := entity.NewRestaurant("Restaurante do João", "13729291719002", "Rua Da Flor 1321")
	suite.Nil(err)

	restaurantRepo := database.NewRestaurantDB(db)
	restaurantRepo.CreateRestaurant(restaurant)

	review, err := entity.NewReview("Restaurante muito bom de extrema qualidade.", 4.94, restaurant.ID)
	suite.Nil(err)

	reviewRepo := database.NewReviewDB(db)
	reviewRepo.CreateReview(review)

	reviewHandler := NewReviewHandler(reviewRepo)
	suite.ReviewHandler = reviewHandler
	suite.Review = review
	suite.Restaurant = restaurant

	r := chi.NewRouter()
	r.Route("/reviews", func(r chi.Router) {
		r.Post("/", reviewHandler.CreateReview)
		r.Get("/{id}", reviewHandler.GetReviewByID)
		r.Delete("/{id}", reviewHandler.DeleteReview)
		r.Get("/", reviewHandler.GetAllReviews)
		r.Put("/{id}", reviewHandler.UpdateReview)
	})

	suite.Router = r
}

func (suite *ReviewHandlersTestSuit) TestCreateReview() {
	reviewJson := fmt.Sprintf(`{"description": "Excelente comida.", "rating": 4.90, "restaurant_id": "%s"}`, suite.Restaurant.ID)
	body := bytes.NewBufferString(reviewJson)

	req := httptest.NewRequest("POST", "/reviews", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.Router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	suite.Equal(http.StatusCreated, res.StatusCode)
}

func (suite *ReviewHandlersTestSuit) TestGetAllReviews() {
	req := httptest.NewRequest("GET", "/reviews", nil)
	w := httptest.NewRecorder()

	suite.Router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	var reviews []map[string]interface{}
	body, err := io.ReadAll(res.Body)
	suite.Nil(err)

	err = json.Unmarshal(body, &reviews)
	suite.Nil(err)

	suite.Equal(http.StatusOK, res.StatusCode)
	suite.Equal(suite.Review.Description, reviews[0]["description"])
	suite.Equal(suite.Review.Rating, reviews[0]["rating"])
	suite.Equal(suite.Review.RestaurantID.String(), reviews[0]["restaurant_id"])
}

func (suite *ReviewHandlersTestSuit) TestGetReviewByID() {
	req := httptest.NewRequest("GET", fmt.Sprintf("/reviews/%s", suite.Review.ID), nil)
	w := httptest.NewRecorder()

	suite.Router.ServeHTTP(w, req)

	var review map[string]interface{}

	res := w.Result()
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	suite.Nil(err)

	err = json.Unmarshal(body, &review)
	suite.Nil(err)

	suite.Equal(http.StatusOK, res.StatusCode)
	suite.Equal(suite.Review.ID.String(), review["id"])
	suite.Equal(suite.Review.Description, review["description"])
	suite.Equal(suite.Review.Rating, review["rating"])
}

func (suite *ReviewHandlersTestSuit) TestDeleteReview() {
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/reviews/%s", suite.Review.ID), nil)
	w := httptest.NewRecorder()

	suite.Router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	suite.Equal(http.StatusNoContent, res.StatusCode)
}

func (suite *ReviewHandlersTestSuit) TestUpdateReview() {
	reviewJson := fmt.Sprintf(`{"description": "Comida excelente, atendimento perfeito", "rating": 4.8, "restaurant_id": "%s"}`, suite.Restaurant.ID)
	body := bytes.NewBufferString(reviewJson)

	req := httptest.NewRequest("PUT", fmt.Sprintf("/reviews/%s", suite.Review.ID), body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.Router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	var review map[string]interface{}
	resBody, err := io.ReadAll(res.Body)
	suite.Nil(err)

	err = json.Unmarshal(resBody, &review)
	suite.Nil(err)

	suite.Equal(http.StatusOK, res.StatusCode)
	suite.Equal(suite.Review.ID.String(), review["id"])
	suite.Equal("Comida excelente, atendimento perfeito", review["description"])
	suite.Equal(4.8, review["rating"])
}

func TestReviewHandlersTestSuit(t *testing.T) {
	suite.Run(t, new(ReviewHandlersTestSuit))
}
