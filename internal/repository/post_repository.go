package repository

import (
	"context"
	"log"

	"github.com/SamuelSchutz13/SocialDev/internal/db"
	"github.com/SamuelSchutz13/SocialDev/internal/entity"
	"github.com/google/uuid"
)

type PostRepository struct {
	queries *db.Queries
}

func NewPostRepository(queries *db.Queries) *PostRepository {
	return &PostRepository{queries: queries}
}

func GetPostRepository(queries *db.Queries) *PostRepository {
	return &PostRepository{queries: queries}
}

func (r *PostRepository) CreatePost(ctx context.Context, params db.CreatePostParams) (entity.PostResponse, error) {
	post, err := r.queries.CreatePost(ctx, params)

	if err != nil {
		log.Printf("%v", err)
	}

	var postResponse entity.PostResponse

	postResponse.PostID = post.PostID.String()
	postResponse.UserID = post.UserID.String()
	postResponse.Title = post.Title
	postResponse.Content = post.Content
	postResponse.Photo = post.Photo.String
	postResponse.Video = post.Video.String
	postResponse.CreatedAt = post.CreatedAt.Time.String()
	postResponse.UpdatedAt = post.UpdatedAt.Time.String()

	return postResponse, err
}

func (r *PostRepository) GetPost(ctx context.Context, postID uuid.UUID) (entity.PostResponse, error) {
	post, err := r.queries.GetPost(ctx, postID)

	if err != nil {
		log.Printf("%v", err)
		return entity.PostResponse{}, err
	}

	var postResponse entity.PostResponse

	postResponse.PostID = post.PostID.String()
	postResponse.UserID = post.UserID.String()
	postResponse.Title = post.Title
	postResponse.Content = post.Content
	postResponse.Photo = post.Photo.String
	postResponse.Video = post.Video.String
	postResponse.CreatedAt = post.CreatedAt.Time.String()
	postResponse.UpdatedAt = post.UpdatedAt.Time.String()

	return postResponse, err
}

func (r *PostRepository) GetAllUserPosts(ctx context.Context, userID uuid.UUID) ([]entity.PostResponse, error) {
	posts, err := r.queries.GetAllUserPosts(ctx, userID)

	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	var postsResponse []entity.PostResponse

	for _, post := range posts {
		var postResponse entity.PostResponse

		postResponse.PostID = post.PostID.String()
		postResponse.UserID = post.UserID.String()
		postResponse.Title = post.Title
		postResponse.Content = post.Content
		postResponse.Photo = post.Photo.String
		postResponse.Video = post.Video.String
		postResponse.CreatedAt = post.CreatedAt.Time.String()
		postResponse.UpdatedAt = post.UpdatedAt.Time.String()

		postsResponse = append(postsResponse, postResponse)
	}

	return postsResponse, err
}

func (r *PostRepository) GetUserPost(ctx context.Context, params db.GetUserPostParams) (entity.PostResponse, error) {
	post, err := r.queries.GetUserPost(ctx, params)

	if err != nil {
		log.Printf("%v", err)
		return entity.PostResponse{}, err
	}

	var postResponse entity.PostResponse

	postResponse.PostID = post.PostID.String()
	postResponse.UserID = post.UserID.String()
	postResponse.Title = post.Title
	postResponse.Content = post.Content
	postResponse.Photo = post.Photo.String
	postResponse.Video = post.Video.String
	postResponse.CreatedAt = post.CreatedAt.Time.String()
	postResponse.UpdatedAt = post.UpdatedAt.Time.String()

	return postResponse, err
}

func (r *PostRepository) GetAllPosts(ctx context.Context) ([]entity.PostResponse, error) {
	posts, err := r.queries.GetAllPosts(ctx)

	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	var postsResponse []entity.PostResponse

	for _, post := range posts {
		var postResponse entity.PostResponse

		postResponse.PostID = post.PostID.String()
		postResponse.UserID = post.UserID.String()
		postResponse.Title = post.Title
		postResponse.Content = post.Content
		postResponse.Photo = post.Photo.String
		postResponse.Video = post.Video.String
		postResponse.CreatedAt = post.CreatedAt.Time.String()
		postResponse.UpdatedAt = post.UpdatedAt.Time.String()

		postsResponse = append(postsResponse, postResponse)
	}

	return postsResponse, err
}

func (r *PostRepository) UpdatePost(ctx context.Context, params db.UpdatePostParams) (entity.PostResponse, error) {
	post, err := r.queries.UpdatePost(ctx, params)

	if err != nil {
		log.Printf("%v", err)
	}

	var postResponse entity.PostResponse

	postResponse.PostID = post.PostID.String()
	postResponse.UserID = post.UserID.String()
	postResponse.Title = post.Title
	postResponse.Content = post.Content
	postResponse.Photo = post.Photo.String
	postResponse.Video = post.Video.String
	postResponse.CreatedAt = post.CreatedAt.Time.String()
	postResponse.UpdatedAt = post.UpdatedAt.Time.String()

	return postResponse, err
}

func (r *PostRepository) DeletePost(ctx context.Context, postID uuid.UUID) error {
	err := r.queries.DeletePost(ctx, postID)

	if err != nil {
		log.Printf("%v", err)
		return err
	}

	return nil
}
