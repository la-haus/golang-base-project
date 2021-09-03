package sample

import (
	_ "github.com/la-haus/sample/controllers/dto"
	"github.com/la-haus/sample/enums"
	"github.com/la-haus/sample/interfaces/controllers"
	"github.com/la-haus/sample/interfaces/services"
	"github.com/la-haus/sample/utils/log"
	"github.com/labstack/echo/v4"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
)

type sampleController struct {
	service services.SampleServiceInterface
}

func NewSampleController(service services.SampleServiceInterface) controllers.SampleControllerInterface {
	return &sampleController{
		service: service,
	}
}

// GetSample godoc
// @Description sample get endpoint
// @ID sample-get
// @Tags Sample Controller
// @Param X-Application-Id header string true "X-Application-Id"
// @Success 200
// @Header 200 {string} Content-Encoding "gzip"
// @Failure 404
// @Router /test [get]
func (r *sampleController) GetSample(c echo.Context) error {
	span, ctx := tracer.StartSpanFromContext(c.Request().Context(), enums.ControllerRequest)
	defer span.Finish()
	span.SetTag(enums.XApplicationID, c.Request().Header.Get(enums.XApplicationID))

	log.WithSpanContext(span.Context()).Infof("sample log %s", "test")
	response, err := r.service.SampleServiceFunction(ctx)
	if err != nil {
		log.WithSpanContext(span.Context()).Error(err.Error())
		tracer.WithError(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, response)
}
