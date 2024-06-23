package service

import (
	"github.com/yohanesmario/online-book-store/internal/service/book"
	"github.com/yohanesmario/online-book-store/internal/service/order"
	"github.com/yohanesmario/online-book-store/internal/service/user"
	"go.uber.org/fx"
)

var ServiceModule = fx.Module("service", fx.Provide(
	user.NewUserService,
	book.NewBookService,
	order.NewOrderService,
))
