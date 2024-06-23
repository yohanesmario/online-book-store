package user

import (
	"time"

	"github.com/yohanesmario/online-book-store/conf"
	"github.com/yohanesmario/online-book-store/internal/domain/model"
	"github.com/yohanesmario/online-book-store/internal/domain/repository"
	"github.com/yohanesmario/online-book-store/internal/service/common/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(name, email, password string) (*model.User, error)
	Login(email, password string) (*session.SessionData, error)
}

type impl struct {
	userRepository repository.IUserRepository[*gorm.DB]
}

func NewUserService(userRepository repository.IUserRepository[*gorm.DB]) UserService {
	return &impl{userRepository: userRepository}
}

func (i *impl) CreateUser(name, email, password string) (*model.User, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	modifier := "CREATE_USER"
	now := time.Now()

	return i.userRepository.CreateUser(model.User{
		Name:      name,
		Email:     email,
		Password:  string(encryptedPassword),
		CreatedBy: &modifier,
		CreatedAt: &now,
		UpdatedBy: &modifier,
		UpdatedAt: &now,
	})
}

func (i *impl) Login(email string, password string) (*session.SessionData, error) {
	user, err := i.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &session.SessionData{
		UserID:    user.ID,
		Email:     user.Email,
		ExpiredAt: time.Now().Add(time.Duration(conf.Configuration.JWTExpiryInSeconds) * time.Second),
	}, nil
}

var _ UserService = (*impl)(nil)
