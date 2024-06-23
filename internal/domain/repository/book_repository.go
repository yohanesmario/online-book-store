package repository

import (
	"github.com/yohanesmario/online-book-store/internal/domain/model"
	"github.com/yohanesmario/online-book-store/internal/domain/repository/base"
)

type IBookRepository[TDBConnection any] interface {
	base.BaseTransactionalRepository[IBookRepository[TDBConnection], TDBConnection]

	GetBooksByLastID(lastID int32, limit int) ([]*model.Book, error)
	GetBooksByIDs(bookIds []int32) ([]*model.Book, error)
	CreateBooks(books []*model.Book) ([]*model.Book, error)
}
