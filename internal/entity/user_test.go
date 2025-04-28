package entity 

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewUser(t *testing.T) {
	user, err := NewUser("Userteste", "teste@gmail.com", "teste234")
	assert.Nil(t, err)
	assert.Equal(t, "Userteste", user.Username)
	assert.Equal(t, "teste@gmail.com", user.Email)
	assert.Equal(t, "teste234", user.Password)
	assert.NotNil(t, user.ID)
}

func TestCreateNewUserWithInvalidPassword(t *testing.T) {
	user, err := NewUser("Userteste", "teste@gmail.com", "teste23")
	assert.Equal(t, ErrInvalidPassword, err)
	assert.EqualError(t, err, "password must have at least 8 characters")
	assert.Nil(t, user)
}

func TestCreateNewUserWithoutEmail(t *testing.T) {
	user, err := NewUser("Userteste", "", "teste234")
	assert.Equal(t, ErrInvalidEmail, err)
	assert.EqualError(t, err, "email is required")
	assert.Nil(t, user)
}

func TestCreateNewUserWithoutUser(t *testing.T) {
	user, err := NewUser("", "teste@gmail.com", "teste234")
	assert.Equal(t, ErrInvalidUsername, err)
	assert.EqualError(t, err, "username is required")
	assert.Nil(t, user)
}