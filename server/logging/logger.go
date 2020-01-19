package logging

import (
	"context"

	"github.com/sirupsen/logrus"
)

// Ctx represents the logger in the context
type Ctx struct{}

// Logger returns the context logger
func Logger(ctx context.Context) *logrus.Logger {
	log, ok := ctx.Value(Ctx{}).(*logrus.Logger)
	if !ok {
		panic("Error: logger not set in context")
	}
	return log
}
