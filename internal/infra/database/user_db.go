package database

import (
	"gorm.io/gorm"
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	pkgEntity "github.com/PedroNunesBH/go-reviews-api/pkg/entity"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{
		DB: db,
	}
}

func (u *UserDB) CreateUser(user *entity.User) error {
	return u.DB.Create(&user).Error
}

func (u *UserDB) GetUserByID(id pkgEntity.ID) (*entity.User, error) {
	var user *entity.User
	result := u.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}