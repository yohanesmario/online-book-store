package repository

import (
	"github.com/yohanesmario/online-book-store/internal/domain/model"
	"github.com/yohanesmario/online-book-store/internal/domain/repository/base"
)

type IUserRepository[TDBConnection any] interface {
	base.BaseTransactionalRepository[IUserRepository[TDBConnection], TDBConnection]

	CreateUser(user model.User) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id int32) (*model.User, error)
}
