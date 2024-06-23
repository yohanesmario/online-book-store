package repository

import (
	"github.com/yohanesmario/online-book-store/internal/impl/gorm/repository/book"
	"github.com/yohanesmario/online-book-store/internal/impl/gorm/repository/order"
	"github.com/yohanesmario/online-book-store/internal/impl/gorm/repository/orderitem"
	"github.com/yohanesmario/online-book-store/internal/impl/gorm/repository/user"
	"go.uber.org/fx"
)

var RepositoryImplModule = fx.Module("repository_impl", fx.Provide(
	book.NewBookRepository,
	user.NewUserRepository,
	order.NewOrderRepository,
	orderitem.NewOrderItemRepository,
))
