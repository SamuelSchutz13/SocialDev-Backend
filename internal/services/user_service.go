package services

import (
	"context"
	"database/sql"
	"log"

	"github.com/SamuelSchutz13/SocialDev/internal/db"
	"github.com/SamuelSchutz13/SocialDev/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func stringToNull(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}

	return sql.NullString{String: s, Valid: true}
}

func GetUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(username, email, password string) (db.User, error) {
	ctx := context.Background()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return db.User{}, err
	}

	userParams := db.CreateUserParams{
		UserID:   uuid.New(),
		Username: username,
		Password: string(passwordHash),
		Email:    email,
	}

	return s.userRepo.CreateUser(ctx, userParams)
}

func (s *UserService) GetUser(userID uuid.UUID) (db.User, error) {
	ctx := context.Background()
	return s.userRepo.GetUser(ctx, userID)
}

func (s *UserService) GetAllUsers() ([]db.User, error) {
	ctx := context.Background()
	return s.userRepo.GetAllUsers(ctx)
}

func (s *UserService) GetUserWithUsername(username string) (db.GetUserWithUsernameRow, error) {
	ctx := context.Background()
	return s.userRepo.GetUserWithUsername(ctx, username)
}

func (s *UserService) UpdateUser(user_id uuid.UUID, username, email, password, avatar, bio, github, linkedin, website string) (db.User, error) {
	ctx := context.Background()

	return s.userRepo.UpdateUser(ctx, db.UpdateUserParams{
		UserID:   user_id,
		Username: username,
		Email:    email,
		Password: password,
		Avatar:   stringToNull(avatar),
		Bio:      stringToNull(bio),
		Github:   stringToNull(github),
		Linkedin: stringToNull(linkedin),
		Website:  stringToNull(website),
	})
}
