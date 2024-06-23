package order

import (
	"github.com/yohanesmario/online-book-store/internal/domain/model"
	"github.com/yohanesmario/online-book-store/internal/domain/repository"
	"github.com/yohanesmario/online-book-store/internal/domain/repository/base"
	"gorm.io/gorm"
)

type OrderService interface {
	GetOrdersByUserIDAndLastID(userID int32, lastID int32, limit int) ([]*model.Order, []*model.OrderItem, []*model.Book, error)
	CreateOrder(order model.Order, orderItems []*model.OrderItem) (*model.Order, []*model.OrderItem, error)
}

type impl struct {
	orderRepository       repository.IOrderRepository[*gorm.DB]
	orderItemRepository   repository.IOrderItemRepository[*gorm.DB]
	bookRepository        repository.IBookRepository[*gorm.DB]
	dbTransactionProvider base.DBTransactionProvider[*gorm.DB]
}

func NewOrderService(
	orderRepository repository.IOrderRepository[*gorm.DB],
	orderItemRepository repository.IOrderItemRepository[*gorm.DB],
	bookRepository repository.IBookRepository[*gorm.DB],
	dbTransactionProvider base.DBTransactionProvider[*gorm.DB],
) OrderService {
	return &impl{
		orderRepository:       orderRepository,
		orderItemRepository:   orderItemRepository,
		bookRepository:        bookRepository,
		dbTransactionProvider: dbTransactionProvider,
	}
}

func (i *impl) CreateOrder(order model.Order, orderItems []*model.OrderItem) (*model.Order, []*model.OrderItem, error) {
	var orderResult *model.Order
	var orderItemsResult []*model.OrderItem
	var err error

	err = i.dbTransactionProvider.Transaction(func(tx base.IDBConnection[*gorm.DB]) error {
		orderResult, err = i.orderRepository.Use(tx).CreateOrder(order)
		if err != nil {
			return err
		}

		bookIDs := make([]int32, 0, len(orderItems))
		for _, item := range orderItems {
			bookIDs = append(bookIDs, *item.BookID)
		}
		books, err := i.bookRepository.Use(tx).GetBooksByIDs(bookIDs)
		if err != nil {
			return err
		}
		bookMap := make(map[int32]*model.Book)
		for _, book := range books {
			bookMap[book.ID] = book
		}

		for i := range orderItems {
			orderItems[i].OrderID = &orderResult.ID
			orderItems[i].Price = &bookMap[*orderItems[i].BookID].Price
		}

		orderItemsResult, err = i.orderItemRepository.Use(tx).CreateOrderItems(orderItems)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	return orderResult, orderItemsResult, nil
}

func (i *impl) GetOrdersByUserIDAndLastID(userID int32, lastID int32, limit int) ([]*model.Order, []*model.OrderItem, []*model.Book, error) {
	orders, err := i.orderRepository.GetOrdersByUserIDAndLastID(userID, lastID, limit)
	if err != nil {
		return nil, nil, nil, err
	}

	orderIDs := make([]int32, len(orders))
	for i, order := range orders {
		orderIDs[i] = order.ID
	}
	orderItems, err := i.orderItemRepository.GetOrderItemsByOrderIDs(orderIDs)
	if err != nil {
		return nil, nil, nil, err
	}

	bookIDs := make([]int32, 0, len(orderItems))
	for _, item := range orderItems {
		bookIDs = append(bookIDs, *item.BookID)
	}
	books, err := i.bookRepository.GetBooksByIDs(bookIDs)
	if err != nil {
		return nil, nil, nil, err
	}

	return orders, orderItems, books, nil
}

var _ OrderService = (*impl)(nil)
