package plugin

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vg006/vtest/internal/types"
	"github.com/vg006/vtest/internal/utils"
)

func Echo(r *echo.Echo, opts types.Options) {
	r.GET(opts.Path, func(c echo.Context) error {
		echoRoutes := r.Routes()
		var routes []types.Route
		for _, r := range echoRoutes {
			if utils.Contains(opts.ExcludePaths, r.Method) {
				continue
			}
			routes = append(routes, types.Route{Method: r.Method, Path: r.Path})
		}
		return c.JSON(http.StatusOK, routes)
	})
}
