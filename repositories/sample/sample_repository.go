package sample

import (
	"context"
	"github.com/la-haus/sample/interfaces/repositories"
)

type sampleRepositoryImpl struct {
}

// NewSampleRepository create new repository
func NewSampleRepository() repositories.SampleRepository {
	return &sampleRepositoryImpl{}
}

func (s sampleRepositoryImpl) SampleRepositoryFunction(c context.Context) (string, error) {
	return "implement me", nil
}
