package routes

import (
	"net/http"

	"github.com/SamuelSchutz13/SocialDev/internal/db"
	"github.com/SamuelSchutz13/SocialDev/internal/handlers"
	"github.com/SamuelSchutz13/SocialDev/internal/middlewares"
	"github.com/SamuelSchutz13/SocialDev/internal/repository"
	"github.com/SamuelSchutz13/SocialDev/internal/services"
)

func SetupPostRoutes(r *http.ServeMux, queries *db.Queries) {
	postRepo := repository.NewPostRepository(queries)
	postService := services.NewPostService(postRepo)
	postHandler := handlers.NewPostHandler(postService)

	r.HandleFunc("POST /posts", middlewares.ProtectedRoutes(postHandler.CreatePostHandler))
	r.HandleFunc("PATCH /posts/{post_id}", middlewares.ProtectedRoutes(postHandler.UpdatePostHandler))
}
