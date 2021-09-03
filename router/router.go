package router

import (
	"github.com/la-haus/sample/config"
	"github.com/la-haus/sample/controllers"
	"github.com/la-haus/sample/enums"
	"github.com/la-haus/sample/interfaces/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
)

type Router struct {
	server *echo.Echo
	sample routes.SampleRouterInterface
}

func NewRouter(
	server *echo.Echo,
	sample routes.SampleRouterInterface,
) *Router {
	return &Router{
		server: server,
		sample: sample,
	}
}

// Init
// @title Swagger Sample Golang API
// @version 1.0
// @description A sample golang API
// @termsOfService http://swagger.io/terms/
// @BasePath /api/sample
func (r *Router) Init() {
	r.server.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))
	apiGroup := r.server.Group(enums.BasePath, middleware.RequestID(), echotrace.Middleware(echotrace.WithServiceName(config.Environments().DDService)))

	apiGroup.GET("/docs/*", EchoWrapHandler())
	apiGroup.GET(enums.HealthPath, controllers.HealthCheck)
	r.sample.Router(apiGroup)
}
