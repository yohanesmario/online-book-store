package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yohanesmario/online-book-store/internal/domain/model"
	"github.com/yohanesmario/online-book-store/internal/service/user"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_CreateUser(t *testing.T) {
	userService, deps := user.NewMockUserService(t)

	deps.UserRepository.EXPECT().
		CreateUser(mock.Anything).
		RunAndReturn(func(u model.User) (*model.User, error) {
			u.ID = 1
			return &u, nil
		})

	user := &model.User{
		Email:    "test@host.co",
		Password: "password",
		Name:     "Test User",
	}

	created, err := userService.CreateUser(user.Name, user.Email, user.Password)
	assert.Nil(t, err)
	assert.NotNil(t, created)
	assert.Equal(t, int32(1), created.ID)
	assert.Equal(t, user.Email, created.Email)
	assert.Equal(t, user.Name, created.Name)

	err = bcrypt.CompareHashAndPassword([]byte(created.Password), []byte(user.Password))
	assert.Nil(t, err)
}

func TestUserService_Login(t *testing.T) {
	userService, deps := user.NewMockUserService(t)

	email := "test@host.co"
	password := "password"
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.Nil(t, err)

	deps.UserRepository.EXPECT().
		GetUserByEmail(mock.Anything).
		RunAndReturn(func(userEmail string) (*model.User, error) {
			return &model.User{
				ID:       1,
				Email:    userEmail,
				Password: string(encryptedPassword),
			}, nil
		})

	sessionData, err := userService.Login(email, password)
	assert.Nil(t, err)
	assert.NotNil(t, sessionData)
	assert.Equal(t, int32(1), sessionData.UserID)
	assert.Equal(t, email, sessionData.Email)
}
