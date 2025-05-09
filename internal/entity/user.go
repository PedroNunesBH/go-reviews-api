package entity 

import (
	"go-reviews-api/pkg/entity"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID  entity.ID    `json:"id"`
	Username  string `json:"username"`
	Email  string    `json:"email"`
	Password string  `json:"password"`
}

var ErrInvalidUsername error = errors.New("username is required")
var ErrInvalidEmail error = errors.New("email is required")
var ErrInvalidPassword error = errors.New("password must have at least 8 characters")

func NewUser(username, email, password string) (*User, error) {
	user := User{
		ID:       entity.NewID(),
		Username: username,
		Email:    email,
		Password: password,
	}

	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hash)
	return &user, nil
}

func (u *User) ValidateUser() error {
	if u.Username == "" {
		return ErrInvalidUsername
	}
	if u.Email == "" {
		return ErrInvalidEmail
	}
	if u.Password == "" || len(u.Password) < 8 {
		return ErrInvalidPassword
	}
	return nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}