package book

import (
	"github.com/yohanesmario/online-book-store/internal/domain/model"
	"github.com/yohanesmario/online-book-store/internal/domain/repository"
	"github.com/yohanesmario/online-book-store/internal/domain/repository/base"
	"gorm.io/gorm"
)

type BookService interface {
	GetBooksByLastID(lastID int32, limit int) ([]*model.Book, error)
	CreateBooks(books []*model.Book) ([]*model.Book, error)
}

type impl struct {
	bookRepository        repository.IBookRepository[*gorm.DB]
	dbTransactionProvider base.DBTransactionProvider[*gorm.DB]
}

func NewBookService(bookRepository repository.IBookRepository[*gorm.DB], dbTransactionProvider base.DBTransactionProvider[*gorm.DB]) BookService {
	return &impl{
		bookRepository:        bookRepository,
		dbTransactionProvider: dbTransactionProvider,
	}
}

func (i *impl) CreateBooks(books []*model.Book) ([]*model.Book, error) {
	var result []*model.Book
	var err error
	err = i.dbTransactionProvider.Transaction(func(txConnection base.IDBConnection[*gorm.DB]) error {
		result, err = i.bookRepository.Use(txConnection).CreateBooks(books)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i *impl) GetBooksByLastID(lastID int32, limit int) ([]*model.Book, error) {
	return i.bookRepository.GetBooksByLastID(lastID, limit)
}

var _ BookService = (*impl)(nil)
