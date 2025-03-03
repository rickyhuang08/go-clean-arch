package sql

import (
	"errors"

	"github.com/rickyhuang08/gin-project/internal/entity"
)

// UserRepository handles DB interactions
type UserRepository struct {
	Users []entity.User // Mock database
}

// NewUserRepository initializes the repo with dummy data
func NewUserRepository() *UserRepository {
	return &UserRepository {
		Users: []entity.User{
			{ID: 1, UserName: "JohnDoe", Email: "john@example.com", Password: "hashed_password"},
		},
	}
}

// FindByEmail fetches a user by email
func (repo *UserRepository) FindByEmail(email string) (*entity.User, error) {
	for _, user := range repo.Users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}