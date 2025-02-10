package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vg006/vtest/internal/types"
	"github.com/vg006/vtest/internal/utils"
)

func HomeHandler(e echo.Context) error {
	return e.String(200, "Hello, World!")
}

func RoutesHandler(e echo.Context) error {
	var req types.ReqSingleUrl
	if err := e.Bind(&req); err != nil {
		fmt.Println("Error binding request:", err)
		return e.String(http.StatusBadRequest, "Invalid request")
	}

	res, err := utils.FetchRoutes(req)
	if err != nil {
		fmt.Println("Error fetching response:", err)
		return e.String(http.StatusInternalServerError, "Error fetching response")
	}
	fmt.Println("Response:", res)
	return e.JSON(http.StatusOK, res)
}

func DetailsHandler(c echo.Context) error {
	var req types.ReqSingleUrl
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	res, err := utils.GetDetails(req)
	if err != nil {
		fmt.Println("Error fetching response:", err)
		return c.String(http.StatusInternalServerError, "Error fetching response")
	}
	fmt.Println("Response:", res)
	return c.JSON(http.StatusOK, res)
}
