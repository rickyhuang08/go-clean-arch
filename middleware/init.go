package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rickyhuang08/gin-project/helpers"
)

type MiddlewareModule struct {
	Time          helpers.TimeProvider
	PublicKeyPath string
}

func NewMiddlewareModule(time helpers.TimeProvider, publicKeyPath string) *MiddlewareModule {
	return &MiddlewareModule{
		Time: time,
	}
}

func (m *MiddlewareModule) InitLoggerMiddleware() *LoggerMiddleware {
	return NewLoggerMiddleWare(m.Time)
}

func (m *MiddlewareModule) InitRateLimiter() *RateLimiter {
	return NewRateLimiter(m.Time)
}

func (m *MiddlewareModule) InitAuthMiddleware() *AuthModule {
	return NewAuthModule(m.Time)
}

func (m *MiddlewareModule) RegisterGlobalMiddleware(r *gin.Engine) {
	r.Use(m.InitLoggerMiddleware().LoggerMiddleware())
	r.Use(CORSMiddleware())
	r.Use(m.InitRateLimiter().RateLimitMiddleware())
}

// RegisterAuthMiddleware applies to protected routes
func (m *MiddlewareModule) RegisterAuthMiddleware(r *gin.RouterGroup) error {
	loadPublicKey, err := LoadPublicKey(m.PublicKeyPath)
	if err != nil {
		return err
	}
	r.Use(m.InitAuthMiddleware().AuthMiddleware(loadPublicKey)) // Requires JWT authentication
	r.Use(m.InitRateLimiter().RateLimitMiddleware())            // Stricter rate limits for authenticated users

	return nil
}
