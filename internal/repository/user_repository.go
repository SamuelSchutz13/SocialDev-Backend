package repository

import (
	"context"
	"log"

	"github.com/SamuelSchutz13/SocialDev/internal/db"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(queries *db.Queries) *UserRepository {
	return &UserRepository{queries: queries}
}

func (r *UserRepository) CreateUser(ctx context.Context, params db.CreateUserParams) (db.User, error) {
	user, err := r.queries.CreateUser(ctx, params)

	if err != nil {
		log.Printf("%v", err)
	}

	return user, err
}
