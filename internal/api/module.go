package api

import (
	"github.com/yohanesmario/online-book-store/internal/api/book"
	"github.com/yohanesmario/online-book-store/internal/api/common/swagger"
	"github.com/yohanesmario/online-book-store/internal/api/healthz"
	"github.com/yohanesmario/online-book-store/internal/api/order"
	"github.com/yohanesmario/online-book-store/internal/api/user"
	"go.uber.org/fx"
)

var APIModule = fx.Module("api", fx.Invoke(
	healthz.RegisterHealthzAPI,
	user.RegisterCreateUserAPI,
	user.RegisterLoginAPI,
	order.RegisterCreateOrderAPI,
	order.RegisterListOrdersAPI,
	book.RegisterListBooksAPI,
	swagger.RegisterSwaggerAPI,
))
