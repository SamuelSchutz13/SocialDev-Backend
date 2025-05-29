package routes

import (
	"net/http"

	"github.com/SamuelSchutz13/SocialDev/internal/db"
)

func SetupRoutes(queries *db.Queries) *http.ServeMux {
	r := http.NewServeMux()

	apiRouter := http.NewServeMux()
	SetupAPIRoutes(apiRouter, queries)

	r.Handle("/api/", http.StripPrefix("/api", apiRouter))

	return r
}
