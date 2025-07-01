package routes

import (
	"net/http"

	"github.com/SamuelSchutz13/SocialDev/internal/db"
)

func SetupAPIRoutes(r *http.ServeMux, queries *db.Queries) {
	SetupUserRoutes(r, queries)
	SetupRoleRoutes(r, queries)
	SetupPostRoutes(r, queries)
}
