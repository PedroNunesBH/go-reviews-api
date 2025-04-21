package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/PedroNunesBH/go-reviews-api/internal/dto"
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	"github.com/PedroNunesBH/go-reviews-api/internal/infra/database"
	"github.com/go-chi/chi"
	pkgEntity "github.com/PedroNunesBH/go-reviews-api/pkg/entity"
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

func (h *ReviewHandler) GetReviewByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsedID, err := pkgEntity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	review, err := h.ReviewRepo.FindReviewByID(parsedID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	reviewJson, err := json.Marshal(&review)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(reviewJson)
}

func (h *ReviewHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsedID, err := pkgEntity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ReviewRepo.DeleteReview(parsedID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *ReviewHandler) GetAllReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := h.ReviewRepo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	reviewsJson, err := json.Marshal(&reviews)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(reviewsJson)
}

func (h *ReviewHandler) UpdateReview(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsedID, err := pkgEntity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	review, err := h.ReviewRepo.FindReviewByID(parsedID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	reviewDTO := &dto.ReviewDTO{}
	err = json.NewDecoder(r.Body).Decode(reviewDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	review.Description = reviewDTO.Description
	review.Rating = reviewDTO.Rating
	review.RestaurantID = reviewDTO.RestaurantID

	err = h.ReviewRepo.UpdateReview(review)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	reviewJson, err := json.Marshal(&review)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(reviewJson)
}