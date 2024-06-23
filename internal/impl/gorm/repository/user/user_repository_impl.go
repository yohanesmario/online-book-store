package user

import (
	"github.com/yohanesmario/online-book-store/internal/domain/model"
	"github.com/yohanesmario/online-book-store/internal/domain/query"
	"github.com/yohanesmario/online-book-store/internal/domain/repository"
	"github.com/yohanesmario/online-book-store/internal/domain/repository/base"
	"gorm.io/gorm"
)

type impl struct {
	db base.IDBConnection[*gorm.DB]
}

func NewUserRepository(db base.IDBConnection[*gorm.DB]) repository.IUserRepository[*gorm.DB] {
	return &impl{db: db}
}

func (i *impl) CreateUser(user model.User) (*model.User, error) {
	u := query.Use(i.db.GetDBConnection()).User
	err := u.Create(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (i *impl) GetUserByEmail(email string) (*model.User, error) {
	u := query.Use(i.db.GetDBConnection()).User
	user, err := u.Where(u.Email.Eq(email)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (i *impl) GetUserByID(id int32) (*model.User, error) {
	u := query.Use(i.db.GetDBConnection()).User
	user, err := u.Where(u.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (i *impl) Use(txConnection base.IDBConnection[*gorm.DB]) repository.IUserRepository[*gorm.DB] {
	return &impl{db: txConnection}
}

var _ repository.IUserRepository[*gorm.DB] = (*impl)(nil)
