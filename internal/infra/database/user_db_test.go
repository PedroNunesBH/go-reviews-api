package database

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	"gorm.io/gorm"
	"github.com/glebarez/sqlite"
	"errors"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}
	db.AutoMigrate(&entity.User{})
	userDB := NewUserDB(db)

	user, err := entity.NewUser("Userteste", "teste@gmail.com", "teste234")
	assert.Nil(t, err)
	err = userDB.CreateUser(user)
	assert.Nil(t, err)
	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)

	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Username, userFound.Username)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
}

func TestGetUserByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})
	userDB := NewUserDB(db)

	user, err := entity.NewUser("Userteste", "teste@gmail.com", "teste234")
	assert.Nil(t, err)
	err = userDB.CreateUser(user)
	assert.Nil(t, err)
	userFound, err := userDB.GetUserByID(user.ID)
	assert.Nil(t, err)

	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Username, user.Username)
	assert.Equal(t, userFound.Email, user.Email)
	assert.Equal(t, userFound.Password, user.Password)
}

func TestGetAllUsers(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})
	userDB := NewUserDB(db)

	user, err := entity.NewUser("Userteste", "teste@gmail.com", "teste234")
	assert.Nil(t, err)
	err = userDB.CreateUser(user)
	assert.Nil(t, err)
	secondUser, err := entity.NewUser("Userteste2", "teste2@gmail.com", "teste23456")
	assert.Nil(t, err)
	err = userDB.CreateUser(secondUser)
	assert.Nil(t, err)
	totalUsers := 2

	users, err := userDB.GetAllUsers()
	assert.Nil(t, err)
	assert.Equal(t, len(users), totalUsers)
}

func TestDeleteUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)
	db.AutoMigrate(&entity.User{})
	userDB := NewUserDB(db)

	user, err := entity.NewUser("Userteste", "teste@gmail.com", "teste234")
	assert.Nil(t, err)
	err = userDB.CreateUser(user)
	assert.Nil(t, err)

	err = userDB.DeleteUser(user.ID)
	assert.Nil(t, err)
	_, err = userDB.GetUserByID(user.ID)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
