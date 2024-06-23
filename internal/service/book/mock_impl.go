package book

import (
	"testing"

	"github.com/yohanesmario/online-book-store/internal/domain/repository"
	"github.com/yohanesmario/online-book-store/internal/domain/repository/base"
	"gorm.io/gorm"
)

type mockDeps struct {
	MockBookRepository        *repository.MockIBookRepository[*gorm.DB]
	MockDBTransactionProvider *base.MockDBTransactionProvider[*gorm.DB]
}

func NewMockBookService(t *testing.T) (BookService, mockDeps) {
	bookRepo := repository.NewMockIBookRepository[*gorm.DB](t)
	dbTransactionProvider := base.NewMockDBTransactionProvider[*gorm.DB](t)
	bookService := NewBookService(bookRepo, dbTransactionProvider)

	return bookService, mockDeps{
		MockBookRepository:        bookRepo,
		MockDBTransactionProvider: dbTransactionProvider,
	}
}
