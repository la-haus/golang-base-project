package services

import (
	"context"
)

type SampleServiceInterface interface {
	SampleServiceFunction(c context.Context) (string, error)
}
