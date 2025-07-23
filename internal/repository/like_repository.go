package repository

import (
	"context"
	"log"

	"github.com/SamuelSchutz13/SocialDev/internal/db"
)

type LikeRepository struct {
	queries *db.Queries
}

func NewLikeRepository(queries *db.Queries) *LikeRepository {
	return &LikeRepository{queries: queries}
}

func GetLikeRepository(queries *db.Queries) *LikeRepository {
	return &LikeRepository{queries: queries}
}

func (r *LikeRepository) CreateLike(ctx context.Context, params db.CreateLikeParams) (db.Like, error) {
	like, err := r.queries.CreateLike(ctx, params)

	if err != nil {
		log.Printf("%v", err)
	}

	return like, err
}

func (r *LikeRepository) GetLike(ctx context.Context, params db.GetLikeParams) (db.Like, error) {
	like, err := r.queries.GetLike(ctx, params)

	if err != nil {
		log.Printf("%v", err)
	}

	return like, err
}

func (r *LikeRepository) DeleteLike(ctx context.Context, params db.DeleteLikeParams) error {
	err := r.queries.DeleteLike(ctx, params)

	if err != nil {
		log.Printf("%v", err)
		return err
	}

	return err
}
