package sample

import (
	"context"
	"github.com/la-haus/sample/enums"
	"github.com/la-haus/sample/interfaces/repositories"
	"github.com/la-haus/sample/interfaces/services"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type sampleService struct {
	sampleRepository repositories.SampleRepository
}

func NewSampleService(sampleRepository repositories.SampleRepository) services.SampleServiceInterface {
	return &sampleService{sampleRepository: sampleRepository}
}

func (e sampleService) SampleServiceFunction(c context.Context) (string, error) {
	span, ctx := tracer.StartSpanFromContext(c, enums.ServiceFunction)
	defer span.Finish()

	response, err := e.sampleRepository.SampleRepositoryFunction(ctx)
	if err != nil {
		return "", err
	}
	return response, nil
}
