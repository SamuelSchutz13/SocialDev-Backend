package routes

import (
	"net/http"

	"github.com/SamuelSchutz13/SocialDev/internal/db"
	"github.com/SamuelSchutz13/SocialDev/internal/handlers"
	"github.com/SamuelSchutz13/SocialDev/internal/repository"
	"github.com/SamuelSchutz13/SocialDev/internal/services"
)

func SetupLikeRoutes(r *http.ServeMux, queries *db.Queries) {
	likeRepo := repository.NewLikeRepository(queries)
	likeService := services.NewLikeService(likeRepo)
	likeHandler := handlers.NewLikeHandler(likeService)

	r.HandleFunc("POST /likes/users/{user_id}/likes/{like_id}", likeHandler.CreateLikeHandler)
	r.HandleFunc("DELETE /likes/users/{user_id}/posts/{post_id}", likeHandler.DeleteLikeHandler)
}
