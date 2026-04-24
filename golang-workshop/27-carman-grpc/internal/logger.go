package internal

import (
	"context"

	"github.com/sirupsen/logrus"
)

func Logger(ctx context.Context) *logrus.Entry {
	logger := logrus.StandardLogger()
	entry := logrus.NewEntry(logger)

	if v, ok := ctx.Value("x-trace-id").(string); ok && v != "" {
		entry = entry.WithField("x-trace-id", v)
	}

	return entry
}
