package log

import (
	"context"
	"github.com/la-haus/sample/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"os"
	"path"
	"runtime"
	"strconv"
	"sync"
)

type traceLog struct {
	Service string `json:"service"`
	Version string `json:"version"`
	Env     string `json:"env"`
	TraceID string `json:"trace_id"`
	SpanID  string `json:"span_id"`
}

const (
	ddField   = "dd"
	fileField = "file"
	lineField = "line"
)

var (
	log   *logrus.Logger
	once  sync.Once
	trace = traceLog{
		Service: config.Environments().DDService,
		Version: config.Environments().DDVersion,
		Env:     config.Environments().DDEnv,
	}
)

func WithContext(c context.Context) *logrus.Entry {
	once.Do(func() {
		log = logrus.New()
		log.Out = os.Stdout
		log.SetFormatter(&logrus.JSONFormatter{})
	})

	_, file, line, _ := runtime.Caller(1)
	return logrus.NewEntry(log).
		WithContext(c).
		WithFields(logrus.Fields{
			ddField: trace,
		}).
		WithFields(logrus.Fields{
			fileField: path.Base(file),
			lineField: line,
		})
}

func WithSpanContext(c ddtrace.SpanContext) *logrus.Entry {
	once.Do(func() {
		log = logrus.New()
		log.Out = os.Stdout
		log.SetFormatter(&logrus.JSONFormatter{})
	})
	_, file, line, _ := runtime.Caller(1)

	trace.TraceID = strconv.FormatUint(c.TraceID(), 10)
	trace.SpanID = strconv.FormatUint(c.SpanID(), 10)

	return logrus.NewEntry(log).
		WithFields(logrus.Fields{
			ddField: trace,
		}).
		WithFields(logrus.Fields{
			fileField: path.Base(file),
			lineField: line,
		})
}
