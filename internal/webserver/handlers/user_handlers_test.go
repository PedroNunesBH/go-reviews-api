package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	"github.com/PedroNunesBH/go-reviews-api/internal/infra/database"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"github.com/stretchr/testify/suite"
	"encoding/json"
	"io"
	"fmt"
	"github.com/go-chi/chi"
	"bytes"
)

type UserHandlersTestSuit struct {
	suite.Suite
	UserHandler *UserHandler
	user *entity.User
}

func (suite *UserHandlersTestSuit) SetupTest() {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	suite.Nil(err)
	db.AutoMigrate(&entity.User{})

	user, err := entity.NewUser("teste132", "teste2gmail.com", "teste234")
	suite.Nil(err)

	userRepo := database.NewUserDB(db)

	userRepo.CreateUser(user)

	userHandler := NewUserHandler(userRepo)

	suite.UserHandler = userHandler
	suite.user = user
}

func (suite *UserHandlersTestSuit) TestGetAllUsers() {
	req := httptest.NewRequest("GET", "/reviews", nil)
	w := httptest.NewRecorder()	
	suite.UserHandler.GetAllUsers(w, req)

	res := w.Result()
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	suite.Nil(err)

	var users []map[string]interface{}

	err = json.Unmarshal(body, &users)
	suite.Require().NoError(err)

	suite.Equal(suite.user.Username, users[0]["username"])
	suite.Equal(suite.user.Email, users[0]["email"])

	suite.Equal(http.StatusOK, res.StatusCode)
}

func (suite *UserHandlersTestSuit) TestGetUser() {
    r := chi.NewRouter()
    
    r.Get("/reviews/{id}", suite.UserHandler.GetUserByID)

    req := httptest.NewRequest("GET", fmt.Sprintf("/reviews/%s", suite.user.ID), nil)
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    res := w.Result()
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    suite.Nil(err)

    var user map[string]interface{}
    err = json.Unmarshal(body, &user)
    suite.Nil(err)

    suite.Equal(suite.user.Username, user["username"])
    suite.Equal(suite.user.Email, user["email"])
}

func (suite *UserHandlersTestSuit) TestCreateUser() {
	userJson := `{"username": "teste234", "email": "teste@gmail.com", "password": "teste234"}`
	body := bytes.NewBufferString(userJson)

	req := httptest.NewRequest("POST", "/reviews", body)
	w := httptest.NewRecorder()
	suite.UserHandler.CreateUser(w, req)

	res := w.Result()
	defer res.Body.Close()

	suite.Equal(201, res.StatusCode)
}

func TestUserHandlersTestSuit(t *testing.T) {
    suite.Run(t, new(UserHandlersTestSuit))
}