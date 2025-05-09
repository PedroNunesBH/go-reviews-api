package database

import (
	"gorm.io/gorm"
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	pkgEntity "go-reviews-api/pkg/entity"
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

func (u *UserDB) GetAllUsers() ([]*entity.User, error) {
	var users []*entity.User
	result := u.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (u *UserDB) DeleteUser(id pkgEntity.ID) error {
	result := u.DB.Where("id = ?", id).Delete(&entity.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserDB) UpdateUser(id pkgEntity.ID, user *entity.User) error {
	result := u.DB.Model(&entity.User{}).
        Where("id = ?", user.ID).
        Updates(map[string]interface{}{
            "username": user.Username,
            "email": user.Email,
            "password": user.Password,
        })

    return result.Error
}