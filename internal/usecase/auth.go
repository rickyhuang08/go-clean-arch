package usecase

import (
	"errors"

	"github.com/rickyhuang08/gin-project/internal/entity"
	"github.com/rickyhuang08/gin-project/internal/repository/sql"
)

// AuthUsecase handles authentication
type AuthUsecase struct {
	UserRepo *sql.UserRepository
}

// NewAuthUsecase initializes auth usecase
func NewAuthUsecase(userRepo *sql.UserRepository) *AuthUsecase {
	return &AuthUsecase{UserRepo: userRepo}
}

// Login checks user credentials (dummy check for now)
func (uc *AuthUsecase) Login(req entity.LoginRequest) (string, error) {
	_, err := uc.UserRepo.FindByEmail(req.Email)
	if err != nil || req.Password != "hashed_password" {
		return "", errors.New("invalid credentials")
	}

	// Mock JWT token
	return "mock.jwt.token", nil
}