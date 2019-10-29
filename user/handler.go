package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler interface {
	Find(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) UserHandler {
	return &userHandler{
		userService,
	}
}

func (h *userHandler) Find(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	user, err := h.userService.Find(params["id"])
	if err != nil {
		http.Error(w, err.Error(), 404)
	}

	response, _ := json.Marshal(user)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)

}

func (h *userHandler) FindAll(w http.ResponseWriter, r *http.Request) {

	users, err := h.userService.FindAll()

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	response, _ := json.Marshal(users)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)

}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 401)
	}

	err = h.userService.Create(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *userHandler) Delete(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	err := h.userService.Delete(params["id"])
	if err != nil {
		http.Error(w, err.Error(), 404)
	}

	w.WriteHeader(http.StatusOK)

}