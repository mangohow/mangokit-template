package middleware

import (
	"context"
	"github.com/mangohow/mangokit/tools"
	"github.com/mangohow/mangokit/transport/http"
	"github.com/sirupsen/logrus"
)

func Logger(logger *logrus.Logger) http.Middleware {
	return func(ctx context.Context, req interface{}, handler http.NextHandler) error {
		c := tools.GinCtxFromContext(ctx)
		logger.Infof("Request Path: %s", c.Request.URL.Path)

		return handler(ctx, req)
	}
}
