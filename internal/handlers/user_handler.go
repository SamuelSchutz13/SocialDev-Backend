package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	configs "github.com/SamuelSchutz13/SocialDev/config"
	"github.com/SamuelSchutz13/SocialDev/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

type UserWithUsername struct {
	Username string
}

type LoginUsderRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Avatar   string    `json:"avatar"`
	Bio      string    `json:"bio"`
	Github   string    `json:"github"`
	Linkedin string    `json:"linkedin"`
	Website  string    `json:"website"`
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
	}

	err = validate.Var(email, "required,email")

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid email: %v", err), http.StatusBadRequest)
	}

	err = validate.Var(password, "required,min=6,max=12")

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid password: %v", err), http.StatusBadRequest)
	}

	_, err = h.userService.CreateUser(username, email, password)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Path
	userID := strings.TrimPrefix(user_id, "/user/")

	uuid, err := uuid.Parse(userID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid user ID format: %v", err), http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUser(uuid)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get user: %v", err), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(user)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (h *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetAllUsers()

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get users: %v", err), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(users)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal users: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (h *UserHandler) GetUserWithUsernameHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	validate = validator.New()

	err := validate.Var(username, "required,max=30")

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid username: %v", err), http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserWithUsername(username)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get user with username: %v", err), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(user)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Path
	userID := strings.TrimPrefix(user_id, "/user/update/")

	convertUUID, err := uuid.Parse(userID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid user ID: %v", err), http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUser(convertUUID)

	if err != nil {
		http.Error(w, fmt.Sprintf("User does not exist: %v", err), http.StatusBadRequest)
		return
	}

	rb := r.Body
	reader, err := io.ReadAll(rb)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get user: %v", err), http.StatusBadRequest)
		return
	}

	rb.Close()

	var userUpdate *UpdateUserRequest
	err = json.Unmarshal(reader, &userUpdate)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error unmarshalling request body: %v", err), http.StatusBadRequest)
		return
	}

	validate = validator.New()

	err = validate.Var(userUpdate.Username, "required,min=3,max=30,alpha")

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid username: %v", err), http.StatusBadRequest)
		return
	}

	err = validate.Var(userUpdate.Email, "required,email")

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid email: %v", err), http.StatusBadRequest)
		return
	}

	err = validate.Var(userUpdate.Password, "required,min=6,max=12")

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid password: %v", err), http.StatusBadRequest)
		return
	}

	if userUpdate.Username == "" {
		userUpdate.Username = user.Username
	}

	if userUpdate.Email == "" {
		userUpdate.Email = user.Email
	}

	if userUpdate.Password == "" {
		userUpdate.Password = user.Password
	}

	updateUser, err := h.userService.UpdateUser(convertUUID, userUpdate.Username,
		userUpdate.Email, userUpdate.Password, userUpdate.Avatar, userUpdate.Bio, userUpdate.Github, userUpdate.Linkedin, userUpdate.Website)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating user: %v", err), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(updateUser)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error to marshal user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Path
	userID := strings.TrimPrefix(user_id, "/user/")

	uuid, err := uuid.Parse(userID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid user ID format: %v", err), http.StatusBadRequest)
		return
	}

	_, err = h.userService.GetUser(uuid)

	if err != nil {
		http.Error(w, fmt.Sprintf("User does not exist: %v", err), http.StatusBadRequest)
		return
	}

	err = h.userService.DeleteUser(uuid)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

func (h *UserHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	rb := r.Body
	reader, err := io.ReadAll(rb)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read request body: %v", err), http.StatusBadRequest)
		return
	}

	rb.Close()

	var userLogin *UpdateUserRequest

	err = json.Unmarshal(reader, &userLogin)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal request body: %v", err), http.StatusBadRequest)
		return
	}

	validate = validator.New()

	err = validate.Var(userLogin.Email, "required,email")

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid email: %v", err), http.StatusBadRequest)
		return
	}

	err = validate.Var(userLogin.Password, "required,min=6,max=12")

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid password: %v", err), http.StatusBadRequest)
		return
	}

	user, err := h.userService.LoginUser(userLogin.Email)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to login user: %v", err), http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))

	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	res, err := json.Marshal(user)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal user: %v", err), http.StatusInternalServerError)
		return
	}

	token, err := configs.CreateToken(user.Username)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to created token: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
	w.Write([]byte(token))
}
