package database

import (
	"gorm.io/gorm"
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
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