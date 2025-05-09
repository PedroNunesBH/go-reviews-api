package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/PedroNunesBH/go-reviews-api/internal/dto"
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	"github.com/PedroNunesBH/go-reviews-api/internal/infra/database"
	"github.com/go-chi/chi"
	pkgEntity "go-reviews-api/pkg/entity"
)

type ReviewHandler struct {
	ReviewRepo *database.ReviewDB
}

func NewReviewHandler(repo *database.ReviewDB) *ReviewHandler {
	return &ReviewHandler{
		ReviewRepo: repo,
	}
}

// Create Review
// @Summary      Create a review
// @Description  Create a review
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Param        request   body     dto.ReviewDTO  true "Review data"
// @Success      201 
// @Failure      400 
// @Failure      500 
// @Router       /reviews/ [post]
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

// Get Review By ID
// @Summary      Get a review by ID
// @Description  Get a review by ID
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Review ID"  Format(uuid)
// @Success      200 
// @Failure      400
// @Failure      404
// @Failure      500 
// @Router       /reviews/{id} [get]
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

// Delete Review
// @Summary      Delete a review 
// @Description  Delete a review
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Review ID"  Format(uuid)
// @Success      204
// @Failure      400
// @Failure      404
// @Router       /reviews/{id} [delete]
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

// Get All Reviews
// @Summary      Get all reviews
// @Description  Get all reviews
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /reviews [get]
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

// Update Review
// @Summary      Update a review
// @Description  Update a review
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Review ID"  Format(uuid)
// @Param        request  body  dto.ReviewDTO  true  "Review Data"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /reviews/{id} [put]
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