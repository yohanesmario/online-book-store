package auth

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/yohanesmario/online-book-store/conf"
	"github.com/yohanesmario/online-book-store/internal/api/common/errorcode"
	"github.com/yohanesmario/online-book-store/internal/api/common/res"
)

var JWTAuth = jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{
		JWTAlg: jwtware.HS512,
		Key:    []byte(conf.Configuration.JWTSecret),
	},
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": res.ErrorData{
				Code:    errorcode.Unauthorized,
				Message: "Unauthorized",
			},
		})
	},
})
