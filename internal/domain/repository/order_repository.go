package repository

import (
	"github.com/yohanesmario/online-book-store/internal/domain/model"
	"github.com/yohanesmario/online-book-store/internal/domain/repository/base"
)

type IOrderRepository[TDBConnection any] interface {
	base.BaseTransactionalRepository[IOrderRepository[TDBConnection], TDBConnection]

	GetOrdersByUserIDAndLastID(userID int32, lastID int32, limit int) ([]*model.Order, error)
	GetOrdersByIDs(orderIDs []int32) ([]*model.Order, error)
	CreateOrder(order model.Order) (*model.Order, error)
}
