package database

import (
	"gorm.io/gorm"
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	pkgEntity "go-reviews-api/pkg/entity"

)

type ReviewDB struct {
	DB *gorm.DB
}

func NewReviewDB(db *gorm.DB) *ReviewDB {
	return &ReviewDB{DB: db}
}

func (r *ReviewDB) FindAll() ([]*entity.Review, error) {
	var reviews []*entity.Review
	result := r.DB.Find(&reviews)
	if result.Error != nil {
		return nil, result.Error
	}
	return reviews, nil
}

func (r *ReviewDB) CreateReview(review *entity.Review) error {
	restaurantFound := &entity.Restaurant{}
	result := r.DB.Where("ID = ?", review.RestaurantID).First(restaurantFound)
	if result.Error != nil {
		return result.Error
	}
	return r.DB.Create(&review).Error
}

func (r *ReviewDB) FindReviewByID(id pkgEntity.ID) (*entity.Review, error) {
	var review *entity.Review
	if err := r.DB.Where("ID = ?", id).First(&review).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func (r *ReviewDB) DeleteReview(id pkgEntity.ID) error {
	result := r.DB.Where("id = ?", id).Delete(&entity.Review{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ReviewDB) UpdateReview(review *entity.Review) error {
	restaurantFound := &entity.Restaurant{}
	result := r.DB.Where("id = ?", review.RestaurantID).First(restaurantFound)
	if result.Error != nil {
		return result.Error
	}
	result = r.DB.Model(&entity.Review{}).
        Where("id = ?", review.ID).
        Updates(map[string]interface{}{
            "description":     review.Description,
            "rating":    review.Rating,
            "restaurant_id": review.RestaurantID,
        })

    return result.Error
}