package user

import (
	"testing"

	"github.com/yohanesmario/online-book-store/internal/domain/repository"
	"gorm.io/gorm"
)

type mockDeps struct {
	UserRepository *repository.MockIUserRepository[*gorm.DB]
}

func NewMockUserService(t *testing.T) (UserService, mockDeps) {
	userRepository := repository.NewMockIUserRepository[*gorm.DB](t)
	userService := NewUserService(userRepository)

	return userService, mockDeps{
		UserRepository: userRepository,
	}
}
