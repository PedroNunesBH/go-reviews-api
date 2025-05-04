package handlers

import (
	"net/http"
	"github.com/PedroNunesBH/go-reviews-api/internal/infra/database"
	"encoding/json"
)

type UserHandler struct {
	UserRepo *database.UserDB
}

func NewUserHandler(repo *database.UserDB) *UserHandler {
	return &UserHandler{
		UserRepo: repo,
	}
}

func (u *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.UserRepo.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	usersJson, err := json.Marshal(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usersJson)
}