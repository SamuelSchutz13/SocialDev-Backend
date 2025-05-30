package handlers

import (
	"fmt"
	"net/http"

	"github.com/SamuelSchutz13/SocialDev/internal/services"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

var validate *validator.Validate

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	validate = validator.New()

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	err := validate.Var(username, "required,min=3,max=30")
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid username: %v", err), http.StatusBadRequest)
		fmt.Println(err)
	}

	err = validate.Var(email, "required,email")
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid email: %v", err), http.StatusBadRequest)
		fmt.Println(err)
	}

	err = validate.Var(password, "required,min=6,max=12")
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid password: %v", err), http.StatusBadRequest)
		fmt.Println(err)
	}

	_, err = h.userService.CreateUser(username, email, password)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
