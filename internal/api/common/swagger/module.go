package swagger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"github.com/yohanesmario/online-book-store/conf"
	_ "github.com/yohanesmario/online-book-store/docs"
)

func RegisterSwaggerAPI(app *fiber.App) {
	if conf.Configuration.SwaggerEnabled {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}
}
