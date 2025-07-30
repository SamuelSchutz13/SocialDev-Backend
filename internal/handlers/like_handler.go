package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SamuelSchutz13/SocialDev/internal/db"
	"github.com/SamuelSchutz13/SocialDev/internal/services"
	"github.com/SamuelSchutz13/SocialDev/utils"
	"github.com/google/uuid"
)

type LikeHandler struct {
	likeService *services.LikeService
}

type LikeResponse struct {
	PostID    uuid.UUID `json:"post_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt string    `json:"created_at"`
}

func NewLikeHandler(likeService *services.LikeService) *LikeHandler {
	return &LikeHandler{likeService: likeService}
}

func (h *LikeHandler) CreateLikeHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("user_id")
	postID := r.PathValue("post_id")

	if userID == "" || postID == "" {
		http.Error(w, "User ID and Like ID are required", http.StatusBadRequest)
		return
	}

	userIdUUID, err := uuid.Parse(userID)

	if err != nil {
		http.Error(w, "Invalid User ID format", http.StatusBadRequest)
		return
	}

	postIdUUID, err := uuid.Parse(postID)

	if err != nil {
		http.Error(w, "Invalid Like ID format", http.StatusBadRequest)
		return
	}

	isLiked, err := h.likeService.GetLike(userIdUUID, postIdUUID)

	if isLiked != (db.Like{}) {
		utils.NewErrorResponse(w, http.StatusBadRequest, "Like already exist", err.Error())
		return
	}

	like, err := h.likeService.CreateLike(userIdUUID, postIdUUID)

	if err != nil {
		http.Error(w, "Failed to create user with like: "+err.Error(), http.StatusInternalServerError)
		return
	}

	likeResponse := LikeResponse{
		PostID:    like.PostID,
		UserID:    like.UserID,
		CreatedAt: like.CreatedAt.Time.String(),
	}

	res, err := json.Marshal(likeResponse)

	if err != nil {
		http.Error(w, "Failed to marshal response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (h *LikeHandler) DeleteLikeHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("user_id")
	postID := r.PathValue("post_id")

	if userID == "" || postID == "" {
		http.Error(w, "User ID and Like ID are required", http.StatusBadRequest)
		return
	}

	userIdUUID, err := uuid.Parse(userID)

	if err != nil {
		http.Error(w, "Invalid User ID format", http.StatusBadRequest)
		return
	}

	postIdUUID, err := uuid.Parse(postID)

	if err != nil {
		http.Error(w, "Invalid Like ID format", http.StatusBadRequest)
		return
	}

	if err := h.likeService.DeleteLike(userIdUUID, postIdUUID); err != nil {
		http.Error(w, "Failed to delete user with like: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils.NewMessageResponse(w, http.StatusOK, "Like deleted successfully")
}
