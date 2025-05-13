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
)

type UserHandlersTestSuit struct {
	suite.Suite
	UserHandler *UserHandler
}

func (suite *UserHandlersTestSuit) SetupTest() {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{})

	userRepo := database.NewUserDB(db)
	userHandler := NewUserHandler(userRepo)

	suite.UserHandler = userHandler
}

func (suite *UserHandlersTestSuit) TestGetAllUsers() {
	req := httptest.NewRequest("GET", "/reviews", nil)
	w := httptest.NewRecorder()	
	suite.UserHandler.GetAllUsers(w, req)

	res := w.Result()
	defer res.Body.Close()

	suite.Equal(http.StatusOK, res.StatusCode)
}

func TestUserHandlersTestSuit(t *testing.T) {
    suite.Run(t, new(UserHandlersTestSuit))
}