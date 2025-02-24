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
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}

type Msg struct {
	Type    MsgType
	Message string
}

type MsgType string

const (
	MsgInfo  MsgType = "INFO"
	MsgError MsgType = "ERROR"
	MsgDone  MsgType = "DONE"
	MsgExit  MsgType = "EXIT"
)

type ReqSingleUrl struct {
	Port     string `json:"port"`
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}

type ResSingleUrl struct {
	StatusCode int      `json:"StatusCode,omitempty"`
	Headers    []Header `json:"Headers"`
}

type Header struct {
	Key      string `json:"Key"`
	Value    string `json:"Value"`
	Usage    string `json:"Usage"`
	Level    string `json:"Level"`
	Presence bool   `json:"Presence"`
}

type SecurityHeaderInfo struct {
	Name       string
	Level      string
	Importance string
}
