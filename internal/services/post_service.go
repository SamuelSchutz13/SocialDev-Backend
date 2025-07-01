package services

import (
	"context"
	"database/sql"
	"log"

	"github.com/SamuelSchutz13/SocialDev/internal/db"
	"github.com/SamuelSchutz13/SocialDev/internal/entity"
	"github.com/SamuelSchutz13/SocialDev/internal/repository"
	"github.com/google/uuid"
)

type PostService struct {
	postRepo *repository.PostRepository
}

func NewPostService(postRepo *repository.PostRepository) *PostService {
	return &PostService{postRepo: postRepo}
}

func (p *PostService) CreatePost(user_id, title, content, photo, video string) (entity.PostResponse, error) {
	ctx := context.Background()

	user_id_UUID, err := uuid.Parse(user_id)

	if err != nil {
		return entity.PostResponse{}, err
	}

	params := db.CreatePostParams{
		UserID:  user_id_UUID,
		Title:   title,
		Content: content,
		Photo:   sql.NullString{String: photo, Valid: photo != ""},
		Video:   sql.NullString{String: video, Valid: video != ""},
	}

	return p.postRepo.CreatePost(ctx, params)
}

func (p *PostService) GetPost(post_id string) (entity.PostResponse, error) {
	ctx := context.Background()

	post_id_UUID, err := uuid.Parse(post_id)

	if err != nil {
		return entity.PostResponse{}, err
	}

	return p.postRepo.GetPost(ctx, post_id_UUID)
}

func (p *PostService) GetAllPosts() ([]entity.PostResponse, error) {
	ctx := context.Background()
	return p.postRepo.GetAllPosts(ctx)
}

func (p *PostService) GetAllUserPosts(user_id string) ([]entity.PostResponse, error) {
	ctx := context.Background()

	user_id_UUID, err := uuid.Parse(user_id)

	if err != nil {
		return nil, err
	}

	return p.postRepo.GetAllUserPosts(ctx, user_id_UUID)
}

func (p *PostService) GetUserPost(post_id, user_id string) (entity.PostResponse, error) {
	ctx := context.Background()

	post_id_UUID, err := uuid.Parse(post_id)

	if err != nil {
		return entity.PostResponse{}, err
	}

	user_id_UUID, err := uuid.Parse(user_id)

	if err != nil {
		return entity.PostResponse{}, err
	}

	return p.postRepo.GetUserPost(ctx, db.GetUserPostParams{PostID: post_id_UUID, UserID: user_id_UUID})
}

func (p *PostService) UpdatePost(post_id, user_id, title, content, photo, video string) (entity.PostResponse, error) {
	ctx := context.Background()

	post_id_UUID, err := uuid.Parse(post_id)

	if err != nil {
		return entity.PostResponse{}, err
	}

	user_id_UUID, err := uuid.Parse(user_id)

	if err != nil {
		return entity.PostResponse{}, err
	}

	_, err = p.postRepo.GetUserPost(ctx, db.GetUserPostParams{PostID: post_id_UUID, UserID: user_id_UUID})

	if err != nil {
		log.Printf("Error getting user post: %v", err)
		return entity.PostResponse{}, err
	}

	params := db.UpdatePostParams{
		PostID:  post_id_UUID,
		UserID:  user_id_UUID,
		Title:   title,
		Content: content,
		Photo:   sql.NullString{String: photo, Valid: photo != ""},
		Video:   sql.NullString{String: video, Valid: video != ""},
	}

	return p.postRepo.UpdatePost(ctx, params)
}

func (p *PostService) DeletePost(post_id, user_id string) error {
	ctx := context.Background()

	post_id_UUID, err := uuid.Parse(post_id)

	if err != nil {
		return err
	}

	user_id_UUID, err := uuid.Parse(user_id)

	if err != nil {
		return err
	}

	_, err = p.postRepo.GetUserPost(ctx, db.GetUserPostParams{PostID: post_id_UUID, UserID: user_id_UUID})

	if err != nil {
		log.Printf("Error getting user post: %v", err)
		return err
	}

	return p.postRepo.DeletePost(ctx, post_id_UUID)
}
