package routes

import (
	"net/http"

	"github.com/SamuelSchutz13/SocialDev/internal/db"
	"github.com/SamuelSchutz13/SocialDev/internal/handlers"
	"github.com/SamuelSchutz13/SocialDev/internal/middlewares"
	"github.com/SamuelSchutz13/SocialDev/internal/repository"
	"github.com/SamuelSchutz13/SocialDev/internal/services"
)

func SetupUserRoutes(r *http.ServeMux, queries *db.Queries) {
	userRepo := repository.NewUserRepository(queries)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r.HandleFunc("POST /users", userHandler.CreateUserHandler)
	r.HandleFunc("GET /user/{user_id}", userHandler.GetUserHandler)
	r.HandleFunc("GET /users/", middlewares.ProtectedRoutes(userHandler.GetAllUsersHandler))
	r.HandleFunc("GET /users/filters", userHandler.GetUserWithUsernameHandler)
	r.HandleFunc("PATCH /user/update/{user_id}", userHandler.UpdateUserHandler)
	r.HandleFunc("DELETE /user/{user_id}", userHandler.DeleteUserHandler)

	r.HandleFunc("POST /user/login", userHandler.LoginUserHandler)
}
