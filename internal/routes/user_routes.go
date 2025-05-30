package routes

import (
	"net/http"

	"github.com/SamuelSchutz13/SocialDev/internal/db"
	"github.com/SamuelSchutz13/SocialDev/internal/handlers"
	"github.com/SamuelSchutz13/SocialDev/internal/repository"
	"github.com/SamuelSchutz13/SocialDev/internal/services"
)

func SetupUserRoutes(r *http.ServeMux, queries *db.Queries) {
	userRepo := repository.NewUserRepository(queries)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r.HandleFunc("POST /user/create", userHandler.CreateUserHandler)
	r.HandleFunc("GET /user/{user_id}", userHandler.GetUserHandler)
	r.HandleFunc("GET /users/", userHandler.GetAllUsersHandler)
}
