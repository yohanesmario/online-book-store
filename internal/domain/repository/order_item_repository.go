package repository

import (
	"github.com/yohanesmario/online-book-store/internal/domain/model"
	"github.com/yohanesmario/online-book-store/internal/domain/repository/base"
)

type IOrderItemRepository[TDBConnection any] interface {
	base.BaseTransactionalRepository[IOrderItemRepository[TDBConnection], TDBConnection]

	GetOrderItemsByOrderIDs(orderIDs []int32) ([]*model.OrderItem, error)
	CreateOrderItems(orderItems []*model.OrderItem) ([]*model.OrderItem, error)
}
