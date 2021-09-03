package sample

import (
	"github.com/la-haus/sample/enums"
	"github.com/la-haus/sample/interfaces/controllers"
	"github.com/la-haus/sample/interfaces/routes"
	"github.com/labstack/echo/v4"
)

type sampleRouter struct {
	controller controllers.SampleControllerInterface
}

func NewSampleRouter(controller controllers.SampleControllerInterface) routes.SampleRouterInterface {
	return &sampleRouter{controller: controller}
}

func (r *sampleRouter) Router(c *echo.Group) {
	c.GET(enums.GetPath, r.controller.GetSample)
}
