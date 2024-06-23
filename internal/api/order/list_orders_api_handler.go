package order

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/yohanesmario/online-book-store/internal/api/common/auth"
	"github.com/yohanesmario/online-book-store/internal/api/common/data"
	"github.com/yohanesmario/online-book-store/internal/api/common/errorcode"
	"github.com/yohanesmario/online-book-store/internal/api/common/res"
	"github.com/yohanesmario/online-book-store/internal/domain/model"
	"github.com/yohanesmario/online-book-store/internal/service/common/session"
	"github.com/yohanesmario/online-book-store/internal/service/order"
)

type listOrdersResponsePayload res.BaseResponse[[]data.OrderData]

// List Orders godoc
//
// @Summary      List Orders
// @Description  List current user's orders, paginated by lastId and limit.
// @Tags         Order APIs
// @Produce      json
//
// @Param  lastId  query  int32  false  "lastId of the fetched orders. Default to 0"
// @Param  limit   query  int    false  "number of orders to fetch. Default to 10"
//
// @Success		200  {object}  listOrdersResponsePayload  "List of orders returned. Will return lastId=-1 if no more orders to fetch"
// @Failure		400  {object}  res.ErrorResponse          "Input can't be parsed"
// @Failure		500  {object}  res.ErrorResponse          "Internal server error"
//
// @Security  BearerAuth
// @Router    /api/v1/orders  [get]
func RegisterListOrdersAPI(app *fiber.App, orderService order.OrderService) {
	app.Get("/api/v1/orders", auth.JWTAuth, func(c *fiber.Ctx) error {
		sessionData, err := session.FromJWTToken(c.Locals("user").(*jwt.Token))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(res.ErrorResponse{
				Error: res.ErrorData{
					Code:    errorcode.InvalidInput,
					Message: "Input can't be parsed",
				},
			})
		}
		lastID := int32(c.QueryInt("lastId", 0))
		limit := c.QueryInt("limit", 10)
		orders, orderItems, books, err := orderService.GetOrdersByUserIDAndLastID(sessionData.UserID, lastID, limit)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(res.ErrorResponse{
				Error: res.ErrorData{
					Code:    errorcode.InvalidInput,
					Message: "Internal server error",
				},
			})
		}
		logrus.Debugf("order size: %d, order item size: %d, book size: %d", len(orders), len(orderItems), len(books))

		orderItemsByOrderIdMap := make(map[int32][]*model.OrderItem)
		for _, item := range orderItems {
			if _, ok := orderItemsByOrderIdMap[*item.OrderID]; !ok {
				orderItemsByOrderIdMap[*item.OrderID] = make([]*model.OrderItem, 0)
			}
			orderItemsByOrderIdMap[*item.OrderID] = append(orderItemsByOrderIdMap[*item.OrderID], item)
		}

		bookByIdMap := make(map[int32]*model.Book)
		for _, book := range books {
			bookByIdMap[book.ID] = book
		}

		result := make([]data.OrderData, len(orders))
		for i := range orders {
			orderItemList := orderItemsByOrderIdMap[orders[i].ID]
			result[i] = data.OrderData{
				ID:     &orders[i].ID,
				UserID: orders[i].UserID,
				Items:  make([]data.OrderItemData, len(orderItemList)),
			}
			for j, item := range orderItemList {
				result[i].Items[j] = data.OrderItemData{
					ID:            &item.ID,
					PurchasePrice: item.Price,
					Quantity:      *item.Quantity,
					Book: data.BookData{
						ID:     bookByIdMap[*item.BookID].ID,
						Title:  bookByIdMap[*item.BookID].Title,
						Author: bookByIdMap[*item.BookID].Author,
						Isbn:   bookByIdMap[*item.BookID].Isbn,
					},
				}
			}

		}

		return c.Status(http.StatusOK).JSON(listOrdersResponsePayload{
			Data: result,
		})
	})
}
