package repositories

import "context"

type SampleRepository interface {
	SampleRepositoryFunction(c context.Context) (string, error)
}
