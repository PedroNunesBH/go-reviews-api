package handlers

import (
	"net/http"
	"github.com/PedroNunesBH/go-reviews-api/internal/infra/database"
	"encoding/json"
	"github.com/PedroNunesBH/go-reviews-api/internal/dto"
	"github.com/PedroNunesBH/go-reviews-api/internal/entity"
	"github.com/go-chi/chi"
	pkgEntity "github.com/PedroNunesBH/go-reviews-api/pkg/entity"
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

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	userDTO := &dto.UserRequestDTO{}
	err := json.NewDecoder(r.Body).Decode(userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := entity.NewUser(userDTO.Username, userDTO.Email, userDTO.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = u.UserRepo.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (u *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsedID, err := pkgEntity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := u.UserRepo.GetUserByID(parsedID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	userResponse := dto.UserResponseDTO{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		
	}
	userJson, err := json.Marshal(&userResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}