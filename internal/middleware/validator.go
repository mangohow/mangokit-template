package middleware

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/mangohow/mangokit-template/api/errors"
	"github.com/mangohow/mangokit/transport/http"
)

var validate = validator.New()

func Validator() http.Middleware {
	return func(ctx context.Context, req interface{}, handler http.Handler) (interface{}, error) {
		err := validate.Struct(req)
		if err != nil {
			return nil, errors.ErrorNameCantNotBeEmpty
		}

		return handler(ctx, req)
	}
}
