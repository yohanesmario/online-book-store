package order

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

func NewOrderRepository(db base.IDBConnection[*gorm.DB]) repository.IOrderRepository[*gorm.DB] {
	return &impl{db: db}
}

func (i *impl) CreateOrder(order model.Order) (*model.Order, error) {
	o := query.Use(i.db.GetDBConnection()).Order
	err := o.Create(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (i *impl) GetOrdersByUserIDAndLastID(userID int32, lastID int32, limit int) ([]*model.Order, error) {
	o := query.Use(i.db.GetDBConnection()).Order
	orders, err := o.Where(o.UserID.Eq(userID), o.ID.Gt(lastID)).
		Order(o.ID.Asc()).
		Limit(limit).
		Find()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (i *impl) GetOrdersByIDs(orderIDs []int32) ([]*model.Order, error) {
	o := query.Use(i.db.GetDBConnection()).Order
	orders, err := o.Where(o.ID.In(orderIDs...)).Find()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (i *impl) Use(txConnection base.IDBConnection[*gorm.DB]) repository.IOrderRepository[*gorm.DB] {
	return &impl{db: txConnection}
}

var _ repository.IOrderRepository[*gorm.DB] = (*impl)(nil)
