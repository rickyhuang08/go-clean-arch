package usecase

import (
	"crypto/rsa"
	"errors"

	"github.com/rickyhuang08/gin-project/internal/entity"
	"github.com/rickyhuang08/gin-project/internal/repository/sql"
	"github.com/rickyhuang08/gin-project/pkg/auth"
)

// AuthUsecase handles authentication
type AuthUsecase struct {
	UserRepo   *sql.UserRepository
	JwtHelper  *auth.JWTHelper
	PrivateKey *rsa.PrivateKey
}

// NewAuthUsecase initializes auth usecase
func NewAuthUsecase(userRepo *sql.UserRepository, jwtHelper *auth.JWTHelper, privateKey *rsa.PrivateKey) *AuthUsecase {
	return &AuthUsecase{
		UserRepo:   userRepo,
		JwtHelper:  jwtHelper,
		PrivateKey: privateKey,
	}
}

// Login checks user credentials (dummy check for now)
func (uc *AuthUsecase) Login(req entity.LoginRequest) (string, error) {
	user, err := uc.UserRepo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if user.Password != req.Password { // Hash comparison needed in real app
		return "", errors.New("invalid credentials")
	}

	token, err := uc.JwtHelper.GenerateJWT(user.ID, uc.PrivateKey)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
