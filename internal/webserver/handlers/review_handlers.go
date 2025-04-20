package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/PedroNunesBH/go-reviews-api/internal/dto"
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	"github.com/PedroNunesBH/go-reviews-api/internal/infra/database"
)

type ReviewHandler struct {
	ReviewRepo *database.ReviewDB
}

func NewReviewHandler(repo *database.ReviewDB) *ReviewHandler {
	return &ReviewHandler{
		ReviewRepo: repo,
	}
}

func (h *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	reviewDTO := &dto.ReviewDTO{}
	err := json.NewDecoder(r.Body).Decode(reviewDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	review, err := entity.NewReview(reviewDTO.Description, reviewDTO.Rating, reviewDTO.RestaurantID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ReviewRepo.CreateReview(review)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}