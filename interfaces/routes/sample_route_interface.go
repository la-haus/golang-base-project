package routes

import "github.com/labstack/echo/v4"

var (
	SampleError = SampleControllerError{Msg: "EXCLUSIVE_CONTROLLER_ERROR"}
)

type SampleControllerError struct {
	Msg string
}

func (e SampleControllerError) Error() string {
	return e.Msg
}

type SampleRouterInterface interface {
	Router(c *echo.Group)
}
