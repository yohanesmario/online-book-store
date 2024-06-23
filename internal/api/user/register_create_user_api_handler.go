package user

import (
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yohanesmario/online-book-store/internal/api/common/errorcode"
	"github.com/yohanesmario/online-book-store/internal/api/common/res"
	"github.com/yohanesmario/online-book-store/internal/api/common/validator"
	"github.com/yohanesmario/online-book-store/internal/service/user"
)

type createUserRequestPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createUserResponsePayload res.BaseResponse[*createUserResponseData]

type createUserResponseData struct {
	UserID    int32      `json:"user_id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
}

// Create User godoc
//
// @Summary      Create User
// @Description  Create user with name, email, and password.
// @Tags         User APIs
// @Accept       json
// @Produce      json
//
// @Param  body  body  createUserRequestPayload  true  "User data"
//
// @Success  201  {object}  createUserResponsePayload  "User created successfully"
// @Failure  400  {object}  res.ErrorResponse          "Input can't be parsed"
// @Failure  500  {object}  res.ErrorResponse          "Internal server error"
//
// @Router  /api/v1/user/register  [post]
func RegisterCreateUserAPI(app *fiber.App, userService user.UserService) {
	app.Post("/api/v1/user/register", func(c *fiber.Ctx) error {
		var payload createUserRequestPayload
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(http.StatusBadRequest).JSON(res.ErrorResponse{
				Error: res.ErrorData{
					Code:    errorcode.InvalidInput,
					Message: "Input can't be parsed",
				},
			})
		}

		normalizedEmail := strings.TrimSpace(strings.ToLower(payload.Email))

		emailValid, err := validator.ValidateEmail(normalizedEmail)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(res.ErrorResponse{
				Error: res.ErrorData{
					Code:    errorcode.UnknownError,
					Message: "Internal server error",
				},
			})
		}
		if !emailValid {
			return c.Status(http.StatusBadRequest).JSON(res.ErrorResponse{
				Error: res.ErrorData{
					Code:    errorcode.InvalidInput,
					Message: "Invalid email format",
				},
			})
		}

		user, err := userService.CreateUser(payload.Name, normalizedEmail, payload.Password)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(res.ErrorResponse{
				Error: res.ErrorData{
					Code:    errorcode.UnknownError,
					Message: "Internal server error",
				},
			})
		}
		user.Password = ""

		return c.Status(http.StatusCreated).JSON(createUserResponsePayload{
			Data: &createUserResponseData{
				UserID:    user.ID,
				Name:      user.Name,
				Email:     user.Email,
				CreatedAt: user.CreatedAt,
			},
		})
	})
}
