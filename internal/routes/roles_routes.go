package routes

import (
	"net/http"

	"github.com/SamuelSchutz13/SocialDev/internal/db"
	"github.com/SamuelSchutz13/SocialDev/internal/handlers"
	"github.com/SamuelSchutz13/SocialDev/internal/repository"
	"github.com/SamuelSchutz13/SocialDev/internal/services"
)

func SetupRoleRoutes(r *http.ServeMux, queries *db.Queries) {
	roleRepo := repository.NewRoleRepository(queries)
	roleService := services.NewRoleService(roleRepo)
	roleHandler := handlers.NewRoleHandler(roleService)

	r.HandleFunc("POST /role/create", roleHandler.CreateRoleHandler)
}
