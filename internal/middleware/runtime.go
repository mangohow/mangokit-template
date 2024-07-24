package middleware

import (
	"context"
	"github.com/mangohow/mangokit/transport/http"
	"github.com/sirupsen/logrus"
	"time"
)

func RunTime(logger *logrus.Logger) http.Middleware {
	return func(ctx context.Context, req interface{}, handler http.NextHandler) error {
		defer runTime(time.Now(), logger)

		return handler(ctx, req)
	}
}

func runTime(start time.Time, logger *logrus.Logger) {
	logger.Infof("run time: %s", time.Now().Sub(start).String())
}
