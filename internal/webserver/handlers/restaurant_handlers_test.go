package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"fmt"

	"testing"

	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	"github.com/PedroNunesBH/go-reviews-api/internal/infra/database"
	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type RestaurantHandlersTestSuit struct {
	suite.Suite
	RestaurantHandler *RestaurantHandler
	Restaurant *entity.Restaurant
	Router http.Handler
}

func (suite *RestaurantHandlersTestSuit) SetupTest() {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	suite.Nil(err)

	db.AutoMigrate(&entity.Restaurant{})

	restaurant, err := entity.NewRestaurant("Restaurante do João", "13729291719002", "Rua Da Flor 1321")
	suite.Nil(err)

	restaurantRepo := database.NewRestaurantDB(db)
	restaurantRepo.CreateRestaurant(restaurant)

	restaurantHandler := NewRestaurantHandler(restaurantRepo)
	suite.RestaurantHandler = restaurantHandler
	suite.Restaurant = restaurant

	r := chi.NewRouter()
	r.Route("/restaurants", func (r chi.Router) {
		r.Post("/", restaurantHandler.CreateRestaurant)
		r.Get("/{id}", restaurantHandler.GetRestaurant)
		r.Delete("/{id}", restaurantHandler.DeleteRestaurant)
		r.Get("/", restaurantHandler.GetAllRestaurants)
		r.Put("/{id}", restaurantHandler.UpdateRestaurant)
	})

	suite.Router = r
}

func (suite *RestaurantHandlersTestSuit) TestCreateRestaurant() {
	restaurantJson := `{"name": "Bar da Esquina", "cnpj": "10293401901925", "address": "Rua 5"}`
	body := bytes.NewBufferString(restaurantJson)

	req := httptest.NewRequest("POST", "/restaurants", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.Router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	suite.Equal(http.StatusCreated, res.StatusCode)
}

func (suite *RestaurantHandlersTestSuit) TestGetAllRestaurants() {
	req := httptest.NewRequest("GET", "/restaurants", nil)
	w := httptest.NewRecorder()

	suite.Router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	var restaurants []map[string]interface{}
	body, err := io.ReadAll(res.Body)
	suite.Nil(err)

	err = json.Unmarshal(body, &restaurants)
	suite.Nil(err)

	suite.Equal(http.StatusOK, res.StatusCode)
	suite.Equal(suite.Restaurant.Name, restaurants[0]["name"])
	suite.Equal(suite.Restaurant.Cnpj, restaurants[0]["cnpj"])
	suite.Equal(suite.Restaurant.Address, restaurants[0]["address"])
}

func (suite *RestaurantHandlersTestSuit) TestGetRestaurantByID() {
	req := httptest.NewRequest("GET", fmt.Sprintf("/restaurants/%s", suite.Restaurant.ID), nil)
	w := httptest.NewRecorder()

	suite.Router.ServeHTTP(w, req)

	var restaurant map[string]interface{}

	res := w.Result()
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	suite.Nil(err)

	err = json.Unmarshal(body, &restaurant)
	suite.Nil(err)

	suite.Equal(http.StatusOK, res.StatusCode)
	suite.Equal(suite.Restaurant.ID.String(), restaurant["id"])
	suite.Equal(suite.Restaurant.Name, restaurant["name"])
}

func (suite *RestaurantHandlersTestSuit) TestDeleteRestaurant() {
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/restaurants/%s", suite.Restaurant.ID), nil)
	w := httptest.NewRecorder()

	suite.Router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	suite.Equal(http.StatusNoContent, res.StatusCode)
}

func (suite *RestaurantHandlersTestSuit) TestUpdateRestaurant() {
	restaurantJson := `{"name": "Restaurante do João", "cnpj": "10293401901925", "address": "Rua 10"}`
	body := bytes.NewBufferString(restaurantJson)

	req := httptest.NewRequest("PUT", fmt.Sprintf("/restaurants/%s", suite.Restaurant.ID), body)
	w := httptest.NewRecorder()

	suite.Router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	var restaurant map[string]interface{}
	resBody, err := io.ReadAll(res.Body)
	suite.Nil(err)

	err = json.Unmarshal(resBody, &restaurant)
	suite.Nil(err)

	suite.Equal(http.StatusOK, res.StatusCode)
	suite.Equal(suite.Restaurant.ID.String(), restaurant["id"])
	suite.Equal(suite.Restaurant.Name, restaurant["name"])
	suite.Equal("Rua 10", restaurant["address"])
}

func TestRestaurantHandlersTestSuit(t *testing.T) {
	suite.Run(t, new(RestaurantHandlersTestSuit))
}
