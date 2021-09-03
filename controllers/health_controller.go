package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Health struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

// HealthCheck godoc
// @Description Check if service is up and healthy
// @Tags Health
// @ID health
// @Success 200
// @Failure 404
// @Router /health [get]
func HealthCheck(c echo.Context) error {
	response := &Health{
		Code:    http.StatusOK,
		Message: "Active!",
	}

	return c.JSON(http.StatusOK, response)
}
