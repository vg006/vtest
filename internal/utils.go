package vtest

import (
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	"github.com/gofiber/fiber"
	"github.com/labstack/echo/v4"
)

func TypeOf(router interface{}) string {
	switch router.(type) {
	case *gin.Engine:
		return "gin"
	case *fiber.App:
		return "fiber"
	case *chi.Mux:
		return "chi"
	case *echo.Echo:
		return "echo"
	default:
		return ""
	}
}
