package orderitem

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

func NewOrderItemRepository(db base.IDBConnection[*gorm.DB]) repository.IOrderItemRepository[*gorm.DB] {
	return &impl{db: db}
}

func (i *impl) CreateOrderItems(orderItems []*model.OrderItem) ([]*model.OrderItem, error) {
	oi := query.Use(i.db.GetDBConnection()).OrderItem
	err := oi.CreateInBatches(orderItems, 1000)
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (i *impl) GetOrderItemsByOrderIDs(orderIDs []int32) ([]*model.OrderItem, error) {
	oi := query.Use(i.db.GetDBConnection()).OrderItem
	orderItems, err := oi.Where(oi.OrderID.In(orderIDs...)).Find()
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (i *impl) Use(txConnection base.IDBConnection[*gorm.DB]) repository.IOrderItemRepository[*gorm.DB] {
	return &impl{db: txConnection}
}

var _ repository.IOrderItemRepository[*gorm.DB] = (*impl)(nil)
