package injector

import (
	"fmt"
	sampleController "github.com/la-haus/sample/controllers/sample"
	sampleRepo "github.com/la-haus/sample/repositories/sample"
	"github.com/la-haus/sample/router"
	sampleRouter "github.com/la-haus/sample/router/sample"
	"github.com/la-haus/sample/server"
	sampleService "github.com/la-haus/sample/services/sample"

	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	checkError(Container.Provide(server.NewServer))
	checkError(Container.Provide(server.NewRequestValidator))
	checkError(Container.Provide(router.NewRouter))
	checkError(Container.Provide(sampleRouter.NewSampleRouter))
	checkError(Container.Provide(sampleService.NewSampleService))
	checkError(Container.Provide(sampleRepo.NewSampleRepository))
	checkError(Container.Provide(sampleController.NewSampleController))

	return Container
}

func checkError(err error) {
	if err != nil {
		panic(fmt.Sprintf("Error injecting %v", err))
	}
}
