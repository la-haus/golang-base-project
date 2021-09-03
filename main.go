package main

import (
	"fmt"
	"github.com/la-haus/sample/server"
	"net"

	"github.com/la-haus/sample/config"
	_ "github.com/la-haus/sample/docs"
	"github.com/la-haus/sample/enums"
	"github.com/la-haus/sample/injector"
	"github.com/la-haus/sample/router"
	"github.com/labstack/echo/v4"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func datadogInitialize() {
	addr := net.JoinHostPort(
		config.Environments().DatadogAgentHost,
		config.Environments().DatadogAgentPort,
	)
	tracer.Start(tracer.WithAgentAddr(addr),
		tracer.WithEnv(config.Environments().DDEnv),
		tracer.WithService(config.Environments().DDService),
		tracer.WithServiceVersion(config.Environments().DDVersion))
}

func main() {
	datadogInitialize()
	defer func() {
		tracer.Stop()
	}()
	container := injector.BuildContainer()
	err := container.Invoke(func(server *echo.Echo, route *router.Router, requestValidator *server.RequestValidator) {

		address := fmt.Sprintf("%s:%s", config.Environments().ServerHost, config.Environments().ServerPort)
		server.Debug = config.Environments().DDEnv == enums.StagingEnvironment
		server.Validator = requestValidator
		route.Init()

		server.Logger.Fatal(server.Start(address))
	})

	if err != nil {
		panic(err)
	}
}
