package server

import "github.com/labstack/echo/v4"

func HomeHandler(e echo.Context) error {
	return e.String(200, "Hello, World!")
}

func RoutesHandler(e echo.Context) error {
	return e.String(200, "Routes")
}
