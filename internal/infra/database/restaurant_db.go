package database

import (
	"gorm.io/gorm"
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	pkgEntity "go-reviews-api/pkg/entity"
)

type RestaurantDB struct {
	DB *gorm.DB
}

func NewRestaurantDB(db *gorm.DB) *RestaurantDB {
	return &RestaurantDB{
		DB: db,
	}
}

func (r *RestaurantDB) CreateRestaurant(restaurant *entity.Restaurant) error {
	return r.DB.Create(&restaurant).Error
}

func (r *RestaurantDB) FindAllRestaurants() ([]*entity.Restaurant, error) {
	var restaurants []*entity.Restaurant
	result := r.DB.Find(&restaurants)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurants, nil
}

func (r *RestaurantDB) FindRestaurantByID(id pkgEntity.ID) (*entity.Restaurant, error) {
	restaurant := &entity.Restaurant{}
	result := r.DB.First(&restaurant, "ID = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurant, nil

}

func (r *RestaurantDB) DeleteRestaurant(id pkgEntity.ID) error {
	result := r.DB.Where("id = ?", id).Delete(&entity.Restaurant{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RestaurantDB) UpdateRestaurant(restaurant *entity.Restaurant) error {
	result := r.DB.Model(&entity.Restaurant{}).
        Where("id = ?", restaurant.ID).
        Updates(map[string]interface{}{
            "name": restaurant.Name,
			"cnpj": restaurant.Cnpj,
			"address": restaurant.Address,
        })

    return result.Error
}