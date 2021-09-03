package controllers

import (
	"github.com/labstack/echo/v4"
)

type SampleControllerInterface interface {
	GetSample(c echo.Context) error
}
