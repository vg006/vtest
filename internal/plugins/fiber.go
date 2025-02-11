package plugin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vg006/vtest/internal/types"
	"github.com/vg006/vtest/internal/utils"
)

func Fiber(app *fiber.App, opts types.Options) {
	app.Get(opts.Path, func(c *fiber.Ctx) error {
		var routes []types.Route
		for _, group := range app.Stack() {
			for _, route := range group {
				// Exclude unwanted methods
				if utils.Contains(opts.ExcludePaths, route.Method) {
					continue
				}
				routes = append(routes, types.Route{
					Method:   route.Method,
					Endpoint: route.Path,
				})
			}
		}
		c.Set("Content-Type", "application/json")
		return c.JSON(routes)
	})
}
