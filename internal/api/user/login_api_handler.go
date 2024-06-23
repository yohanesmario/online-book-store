package user

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yohanesmario/online-book-store/internal/api/common/errorcode"
	"github.com/yohanesmario/online-book-store/internal/api/common/res"
	"github.com/yohanesmario/online-book-store/internal/api/common/validator"
	"github.com/yohanesmario/online-book-store/internal/service/user"
)

type loginRequestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponsePayload res.BaseResponse[*loginResponseData]

type loginResponseData struct {
	JWTToken string `json:"token"`
}

// Login godoc
//
// @Summary      Login
// @Description  Login using email and password, return jwt token.
// @Tags         User APIs
// @Accept       json
// @Produce      json
//
// @Param  body  body  loginRequestPayload  true  "User credentials"
//
// @Success  200  {object}  loginResponsePayload  "Login success, jwt token returned"
// @Failure  400  {object}  res.ErrorResponse     "Input can't be parsed"
// @Failure  500  {object}  res.ErrorResponse     "Internal server error"
//
// @Router  /api/v1/user/login  [post]
func RegisterLoginAPI(app *fiber.App, userService user.UserService) {
	app.Post("/api/v1/user/login", func(c *fiber.Ctx) error {
		req := new(loginRequestPayload)
		if err := c.BodyParser(req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(res.ErrorResponse{
				Error: res.ErrorData{
					Code:    errorcode.InvalidInput,
					Message: "Input can't be parsed",
				},
			})
		}

		normalizedEmail := strings.TrimSpace(strings.ToLower(req.Email))

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

		sessionData, err := userService.Login(normalizedEmail, req.Password)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(res.ErrorResponse{
				Error: res.ErrorData{
					Code:    errorcode.UnknownError,
					Message: "Internal server error",
				},
			})
		}

		jwtToken, err := sessionData.ToJWTString()
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(res.ErrorResponse{
				Error: res.ErrorData{
					Code:    errorcode.InvalidInput,
					Message: "Internal server error",
				},
			})
		}

		return c.Status(http.StatusOK).JSON(loginResponsePayload{
			Data: &loginResponseData{
				JWTToken: jwtToken,
			},
		})
	})
}
