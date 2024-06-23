package order

import (
	"testing"

	"github.com/yohanesmario/online-book-store/internal/domain/repository"
	"github.com/yohanesmario/online-book-store/internal/domain/repository/base"
	"gorm.io/gorm"
)

type mockDeps struct {
	MockOrderRepository       *repository.MockIOrderRepository[*gorm.DB]
	MockOrderItemRepository   *repository.MockIOrderItemRepository[*gorm.DB]
	MockBookRepository        *repository.MockIBookRepository[*gorm.DB]
	MockDBTransactionProvider *base.MockDBTransactionProvider[*gorm.DB]
}

func NewMockOrderService(t *testing.T) (OrderService, mockDeps) {
	orderRepository := repository.NewMockIOrderRepository[*gorm.DB](t)
	orderItemRepository := repository.NewMockIOrderItemRepository[*gorm.DB](t)
	bookRepository := repository.NewMockIBookRepository[*gorm.DB](t)
	dbTransactionProvider := base.NewMockDBTransactionProvider[*gorm.DB](t)
	orderService := NewOrderService(orderRepository, orderItemRepository, bookRepository, dbTransactionProvider)

	return orderService, mockDeps{
		MockOrderRepository:       orderRepository,
		MockOrderItemRepository:   orderItemRepository,
		MockBookRepository:        bookRepository,
		MockDBTransactionProvider: dbTransactionProvider,
	}
}
