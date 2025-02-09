package types

import (
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
)

type (
	Gin   = *gin.Engine
	Fiber = *fiber.App
	Chi   = *chi.Mux
	Echo  = *echo.Echo
)

type Options struct {
	Path           string
	Disable        bool
	ExcludeMethods []string
	ExcludePaths   []string
}

type Route struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}
