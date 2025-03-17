package http

import "github.com/rickyhuang08/gin-project/internal/usecase"

// Handler struct contains all use cases
type Handler struct {
	AuthUsecase *usecase.AuthUsecase
	UserUsecase *usecase.UserUsecase
}

// NewHandler initializes an HTTP handler
func NewHandler(authUC *usecase.AuthUsecase, userUC *usecase.UserUsecase) *Handler {
	return &Handler{AuthUsecase: authUC, UserUsecase: userUC}
}