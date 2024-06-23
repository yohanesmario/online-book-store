package order_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yohanesmario/online-book-store/internal/domain/model"
	"github.com/yohanesmario/online-book-store/internal/domain/repository/base"
	"github.com/yohanesmario/online-book-store/internal/service/order"
	"gorm.io/gorm"
)

func TestOrderService_CreateOrder(t *testing.T) {
	orderService, deps := order.NewMockOrderService(t)

	orderIdCounter := 0
	orderItemIdCounter := 0
	deps.MockDBTransactionProvider.EXPECT().
		Transaction(mock.Anything).
		RunAndReturn(func(txBody func(base.IDBConnection[*gorm.DB]) error) error {
			return txBody(base.NewMockIDBConnection[*gorm.DB](t))
		})
	deps.MockOrderRepository.EXPECT().
		Use(mock.Anything).
		Return(deps.MockOrderRepository)
	deps.MockOrderItemRepository.EXPECT().
		Use(mock.Anything).
		Return(deps.MockOrderItemRepository)
	deps.MockBookRepository.EXPECT().
		Use(mock.Anything).
		Return(deps.MockBookRepository)
	deps.MockOrderRepository.EXPECT().
		CreateOrder(mock.Anything).
		RunAndReturn(func(o model.Order) (*model.Order, error) {
			orderIdCounter += 1
			o.ID = int32(orderIdCounter)
			return &o, nil
		})
	deps.MockOrderItemRepository.EXPECT().
		CreateOrderItems(mock.Anything).
		RunAndReturn(func(oi []*model.OrderItem) ([]*model.OrderItem, error) {
			for i := range oi {
				orderItemIdCounter += 1
				oi[i].ID = int32(orderItemIdCounter)
			}
			return oi, nil
		})
	deps.MockBookRepository.EXPECT().
		GetBooksByIDs(mock.Anything).
		RunAndReturn(func(ids []int32) ([]*model.Book, error) {
			books := make([]*model.Book, len(ids))
			for i := range ids {
				books[i] = &model.Book{ID: ids[i], Price: 1000}
			}
			return books, nil
		})

	items := []*model.OrderItem{}
	for i := 0; i < 3; i++ {
		id := int32(i + 1)
		quantity := int32(10)
		items = append(items, &model.OrderItem{
			BookID:   &id,
			Quantity: &quantity,
		})
	}
	orders, orderItems, err := orderService.CreateOrder(model.Order{}, items)
	assert.Nil(t, err)
	assert.NotNil(t, orders)
	assert.NotNil(t, orderItems)
	assert.Len(t, orderItems, len(items))
	for i := range items {
		assert.Equal(t, items[i].BookID, orderItems[i].BookID)
		assert.Equal(t, items[i].Quantity, orderItems[i].Quantity)
		assert.Equal(t, float64(1000), *orderItems[i].Price)
	}
}

func TestOrderService_GetOrdersByUserIDAndLastID(t *testing.T) {
	orderService, deps := order.NewMockOrderService(t)

	deps.MockOrderRepository.EXPECT().
		GetOrdersByUserIDAndLastID(mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(userId, lastId int32, limit int) ([]*model.Order, error) {
			orders := make([]*model.Order, limit)
			for i := 0; i < limit; i++ {
				orders[i] = &model.Order{ID: int32(int(lastId) + i + 1), UserID: userId}
			}
			return orders, nil
		})
	deps.MockOrderItemRepository.EXPECT().
		GetOrderItemsByOrderIDs(mock.Anything).
		RunAndReturn(func(orderIds []int32) ([]*model.OrderItem, error) {
			orderItems := make([]*model.OrderItem, len(orderIds))
			for i, orderId := range orderIds {
				orderItemId := int32(i + 1)
				orderItems[i] = &model.OrderItem{
					ID:      orderItemId,
					OrderID: &orderId,
					BookID:  &orderItemId,
				}
			}
			return orderItems, nil
		})
	deps.MockBookRepository.EXPECT().
		GetBooksByIDs(mock.Anything).
		RunAndReturn(func(ids []int32) ([]*model.Book, error) {
			books := make([]*model.Book, len(ids))
			for i := range ids {
				books[i] = &model.Book{ID: ids[i]}
			}
			return books, nil
		})

	orders, orderItems, books, err := orderService.GetOrdersByUserIDAndLastID(1, 0, 10)
	assert.Nil(t, err)
	assert.Len(t, orders, 10)
	assert.Len(t, orderItems, 10)
	assert.Len(t, books, 10)
	for i := range orders {
		assert.Equal(t, int32(i+1), orders[i].ID)
		assert.Equal(t, int32(i+1), orderItems[i].ID)
		assert.Equal(t, int32(i+1), books[i].ID)
	}
}
