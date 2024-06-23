package healthz

import (
	"github.com/gofiber/fiber/v2"
)

// Healthz godoc
//
// @Summary     Healthcheck
// @Description Healthcheck api, always expected to return "OK" and http status code 200. If not, something is wrong.
// @Tags        Z-Pages
// @Produce     text/plain
//
// @Success  200  string  string  "OK"
//
// @Router  /healthz  [get]
func RegisterHealthzAPI(app *fiber.App) {
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
