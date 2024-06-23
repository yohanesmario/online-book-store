package order

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yohanesmario/online-book-store/internal/api/common/auth"
	"github.com/yohanesmario/online-book-store/internal/api/common/data"
	"github.com/yohanesmario/online-book-store/internal/api/common/errorcode"
	"github.com/yohanesmario/online-book-store/internal/api/common/res"
	"github.com/yohanesmario/online-book-store/internal/domain/model"
	"github.com/yohanesmario/online-book-store/internal/service/common/session"
	"github.com/yohanesmario/online-book-store/internal/service/order"
)

type createOrderRequestPayload struct {
	Items []data.CreateOrderItemData `json:"items"`
}

type createOrderResponsePayload res.BaseResponse[*createOrderResponseData]

type createOrderResponseData struct {
	OrderID int32 `json:"order_id"`
}

// Create Order godoc
//
// @Summary     Create Order
// @Description Create a new order with list of items which contains bookId and quantity.
// @Tags        Order APIs
// @Accept      json
// @Produce     json
//
// @Param  body  body  createOrderRequestPayload  true  "Order data"
//
// @Success  200  {object}  createOrderResponsePayload  "Create order success, order id returned"
// @Failure  401  {object}  res.ErrorResponse           "Unauthorized"
// @Failure  400  {object}  res.ErrorResponse           "Input can't be parsed"
// @Failure  500  {object}  res.ErrorResponse           "Internal server error"
//
// @Security  BearerAuth
// @Router    /api/v1/order/create  [post]
func RegisterCreateOrderAPI(app *fiber.App, orderService order.OrderService) {
	app.Post("/api/v1/order/create", auth.JWTAuth, func(c *fiber.Ctx) error {
		var payload createOrderRequestPayload
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(res.ErrorResponse{
				Error: res.ErrorData{
					Code:    errorcode.InvalidInput,
					Message: "Input can't be parsed",
				},
			})
		}

		sessionData, err := session.FromJWTToken(c.Locals("user").(*jwt.Token))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(res.ErrorResponse{
				Error: res.ErrorData{
					Code:    errorcode.InvalidInput,
					Message: "Input can't be parsed",
				},
			})
		}

		modifier := strconv.Itoa(int(sessionData.UserID))
		now := time.Now()

		orderItems := make([]*model.OrderItem, len(payload.Items))
		for i, item := range payload.Items {
			orderItems[i] = &model.OrderItem{
				BookID:    &item.BookID,
				Quantity:  &item.Quantity,
				CreatedBy: &modifier,
				CreatedAt: &now,
				UpdatedBy: &modifier,
				UpdatedAt: &now,
			}
		}

		orderResult, _, err := orderService.CreateOrder(
			model.Order{
				UserID:    sessionData.UserID,
				CreatedBy: &modifier,
				CreatedAt: &now,
				UpdatedBy: &modifier,
				UpdatedAt: &now,
			},
			orderItems,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(res.ErrorResponse{
				Error: res.ErrorData{
					Code:    errorcode.InvalidInput,
					Message: "Internal server error",
				},
			})
		}

		return c.Status(fiber.StatusCreated).JSON(createOrderResponsePayload{
			Data: &createOrderResponseData{
				OrderID: orderResult.ID,
			},
		})
	})
}
