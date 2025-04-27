package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/PedroNunesBH/go-reviews-api/internal/dto"
	"github.com/PedroNunesBH/go-reviews-api/internal/infra/database"
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	"github.com/go-chi/chi"
	pkgEntity "github.com/PedroNunesBH/go-reviews-api/pkg/entity"
)

type RestaurantHandler struct {
	RestaurantRepo *database.RestaurantDB
}

func NewRestaurantHandler(repo *database.RestaurantDB) *RestaurantHandler {
	return &RestaurantHandler{
		RestaurantRepo: repo,
	}
}

// Create Restaurant
// @Summary      Creat a restaurant
// @Description  Creat a restaurant
// @Tags         restaurants
// @Accept       json
// @Produce      json
// @Param        request   body     dto.RestaurantDTO  true "Restaurant data"
// @Success      201 
// @Failure      400 
// @Failure      500 
// @Router       /restaurants/ [post]
func (h *RestaurantHandler) CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurantDTO := &dto.RestaurantDTO{}
	err := json.NewDecoder(r.Body).Decode(restaurantDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	restaurant, err := entity.NewRestaurant(restaurantDTO.Name, restaurantDTO.Cnpj, restaurantDTO.Address)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.RestaurantRepo.CreateRestaurant(restaurant)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// Get Restaurant
// @Summary      List restaurants
// @Description  Get a restaurant
// @Tags         restaurants
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "restaurant ID" Format(uuid)
// @Success      200 {array} entity.Restaurant
// @Failure      400 
// @Failure      404
// @Router       /restaurants/{id} [get]
func (h *RestaurantHandler) GetRestaurant(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsedID, err := pkgEntity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	restaurant, err := h.RestaurantRepo.FindRestaurantByID(parsedID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	restaurantJson, err := json.Marshal(&restaurant)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(restaurantJson)
}

// Delete Restaurant
// @Summary      Delete a restaurant
// @Description  Delete a restaurant
// @Tags         restaurants
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "restaurant ID" Format(uuid)
// @Success      204
// @Failure      400 
// @Failure      404
// @Router       /restaurants/{id} [delete]
func (h *RestaurantHandler) DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsedID, err := pkgEntity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.RestaurantRepo.DeleteRestaurant(parsedID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// Get all restaurants
// @Summary      Get all restaurants
// @Description  Get all restaurants
// @Tags         restaurants
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /restaurants [get]
func (h *RestaurantHandler) GetAllRestaurants(w http.ResponseWriter, r *http.Request) {
	restaurants, err := h.RestaurantRepo.FindAllRestaurants()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	restaurantsJson, err := json.Marshal(&restaurants)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(restaurantsJson)
}

// Update a Restaurant
// @Summary      Update a restaurant
// @Description  Update a restaurant
// @Tags         restaurants
// @Accept       json
// @Produce      json
// @Param        id  path  string  true "restaurant ID" Format(uuid)
// @Param        request body dto.RestaurantDTO true "Restaurant data"
// @Success      204
// @Failure      400 
// @Failure      404
// @Router       /restaurants/{id} [put]
func (h *RestaurantHandler) UpdateRestaurant(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsedID, err := pkgEntity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	restaurant, err := h.RestaurantRepo.FindRestaurantByID(parsedID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	restaurantDTO := &dto.RestaurantDTO{}
	err = json.NewDecoder(r.Body).Decode(&restaurantDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	restaurant.Name = restaurantDTO.Name
	restaurant.Cnpj = restaurantDTO.Cnpj
	restaurant.Address = restaurantDTO.Address
	err = h.RestaurantRepo.UpdateRestaurant(restaurant)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	restaurantJson, err := json.Marshal(&restaurant)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(restaurantJson)
}