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

func (h *RestaurantHandler) CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurantDTO := &dto.RestaurantDTO{}
	err := json.NewDecoder(r.Body).Decode(restaurantDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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