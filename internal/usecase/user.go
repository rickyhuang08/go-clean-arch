package usecase

import (
	"github.com/rickyhuang08/gin-project/internal/entity"
	"github.com/rickyhuang08/gin-project/internal/repository/sql"
)

// UserUsecase handles user logic
type UserUsecase struct {
	UserRepo *sql.UserRepository
}

// NewUserUsecase initializes user usecase
func NewUserUsecase(userRepo *sql.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo: userRepo}
}

// GetUserProfile returns mock user profile
func (uc *UserUsecase) GetUserProfile(userID int) (*entity.User, error) {
	for _, user := range uc.UserRepo.Users {
		if user.ID == userID {
			return &user, nil
		}
	}
	return nil, nil
}