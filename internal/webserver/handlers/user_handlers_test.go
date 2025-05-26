package handlers

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/PedroNunesBH/go-reviews-api/internal/entity"
    "github.com/PedroNunesBH/go-reviews-api/internal/infra/database"
    "github.com/go-chi/chi"
    "github.com/glebarez/sqlite"
    "github.com/stretchr/testify/suite"
    "gorm.io/gorm"
)

type UserHandlersTestSuit struct {
    suite.Suite
    UserHandler *UserHandler
    user        *entity.User
    router      http.Handler
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

    r := chi.NewRouter()
    r.Route("/users", func(r chi.Router) {
        r.Get("/", userHandler.GetAllUsers)
        r.Post("/", userHandler.CreateUser)
        r.Get("/{id}", userHandler.GetUserByID)
        r.Delete("/{id}", userHandler.DeleteUser)
        r.Put("/{id}", userHandler.UpdateUser)
    })

    suite.router = r
}

func (suite *UserHandlersTestSuit) TestGetAllUsers() {
    req := httptest.NewRequest("GET", "/users", nil)
    w := httptest.NewRecorder()

    suite.router.ServeHTTP(w, req)

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
    req := httptest.NewRequest("GET", fmt.Sprintf("/users/%s", suite.user.ID), nil)
    w := httptest.NewRecorder()

    suite.router.ServeHTTP(w, req)

    res := w.Result()
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    suite.Nil(err)

    var user map[string]interface{}
    err = json.Unmarshal(body, &user)
    suite.Nil(err)

    suite.Equal(suite.user.Username, user["username"])
    suite.Equal(suite.user.Email, user["email"])
    suite.Equal(http.StatusOK, res.StatusCode)
}

func (suite *UserHandlersTestSuit) TestCreateUser() {
    userJson := `{"username": "teste234", "email": "teste@gmail.com", "password": "teste234"}`
    body := bytes.NewBufferString(userJson)

    req := httptest.NewRequest("POST", "/users", body)
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    suite.router.ServeHTTP(w, req)

    res := w.Result()
    defer res.Body.Close()

    suite.Equal(http.StatusCreated, res.StatusCode)
}

func (suite *UserHandlersTestSuit) TestUpdateUser() {
	userJson := `{"username": "teste234", "email": "teste@gmail.com", "password": "teste234"}`
    body := bytes.NewBufferString(userJson)

	req := httptest.NewRequest("PUT", fmt.Sprintf("/users/%s", suite.user.ID), body)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	suite.Nil(err)

	var user *entity.User
	err = json.Unmarshal(responseBody, &user)
	suite.Nil(err)

	suite.Equal(http.StatusOK, res.StatusCode)
	suite.Equal("teste234", user.Username)
	suite.Equal("teste@gmail.com", user.Email)
	suite.Equal(suite.user.ID, user.ID)

	updatedUser, err := suite.UserHandler.UserRepo.GetUserByID(suite.user.ID)
    suite.Nil(err)
    suite.Equal("teste234", updatedUser.Username)
}

func (suite *UserHandlersTestSuit) TestDeleteUser() {
    req := httptest.NewRequest("DELETE", fmt.Sprintf("/users/%s", suite.user.ID), nil)
    w := httptest.NewRecorder()

    suite.router.ServeHTTP(w, req)

    res := w.Result()
    defer res.Body.Close()

    suite.Equal(http.StatusNoContent, res.StatusCode)
}

func TestUserHandlersTestSuit(t *testing.T) {
    suite.Run(t, new(UserHandlersTestSuit))
}