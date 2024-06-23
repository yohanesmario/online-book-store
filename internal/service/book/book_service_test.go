package book_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yohanesmario/online-book-store/internal/domain/model"
	"github.com/yohanesmario/online-book-store/internal/domain/repository/base"
	"github.com/yohanesmario/online-book-store/internal/service/book"
	"gorm.io/gorm"
)

func TestBookService_CreateBook(t *testing.T) {
	bookService, deps := book.NewMockBookService(t)

	deps.MockDBTransactionProvider.EXPECT().
		Transaction(mock.Anything).
		RunAndReturn(func(txBody func(base.IDBConnection[*gorm.DB]) error) error {
			return txBody(base.NewMockIDBConnection[*gorm.DB](t))
		})
	deps.MockBookRepository.EXPECT().
		Use(mock.Anything).
		Return(deps.MockBookRepository)
	deps.MockBookRepository.EXPECT().
		CreateBooks(mock.Anything).
		RunAndReturn(func(b []*model.Book) ([]*model.Book, error) {
			return b, nil
		})

	source := []*model.Book{
		{
			Title: "Book 1",
		},
		{
			Title: "Book 2",
		},
	}
	res, err := bookService.CreateBooks(source)
	assert.Nil(t, err)
	assert.Len(t, res, len(source))
	for i := range source {
		assert.Equal(t, source[i], res[i])
	}
}

func TestBookService_GetBooksByLastID(t *testing.T) {
	bookService, deps := book.NewMockBookService(t)

	deps.MockBookRepository.EXPECT().
		GetBooksByLastID(mock.Anything, mock.Anything).
		RunAndReturn(func(id int32, limit int) ([]*model.Book, error) {
			mockBooks := []*model.Book{}
			for i := int(id + 1); i <= limit; i++ {
				mockBooks = append(mockBooks, &model.Book{ID: int32(i)})
			}
			return mockBooks, nil
		})

	res, err := bookService.GetBooksByLastID(0, 10)
	assert.Nil(t, err)
	assert.Len(t, res, 10)
	minID := math.MaxInt
	for i := range res {
		if int(res[i].ID) < minID {
			minID = int(res[i].ID)
		}
		assert.Equal(t, int32(i+1), res[i].ID)
	}
	assert.Greater(t, minID, 0)
}
