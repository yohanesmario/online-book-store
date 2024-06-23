package book

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

func NewBookRepository(db base.IDBConnection[*gorm.DB]) repository.IBookRepository[*gorm.DB] {
	return &impl{db: db}
}

func (i *impl) CreateBooks(books []*model.Book) ([]*model.Book, error) {
	b := query.Use(i.db.GetDBConnection()).Book
	err := b.CreateInBatches(books, 1000)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (i *impl) GetBooksByLastID(lastID int32, limit int) ([]*model.Book, error) {
	b := query.Use(i.db.GetDBConnection()).Book
	books, err := b.Where(b.ID.Gt(lastID)).Order(b.ID.Asc()).Limit(limit).Find()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (i *impl) GetBooksByIDs(bookIds []int32) ([]*model.Book, error) {
	b := query.Use(i.db.GetDBConnection()).Book
	books, err := b.Where(b.ID.In(bookIds...)).Find()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (i *impl) Use(txConnection base.IDBConnection[*gorm.DB]) repository.IBookRepository[*gorm.DB] {
	return &impl{db: txConnection}
}

var _ repository.IBookRepository[*gorm.DB] = (*impl)(nil)
