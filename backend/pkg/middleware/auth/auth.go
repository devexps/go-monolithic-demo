package auth

import (
	"context"
	"github.com/devexps/go-micro/v2/middleware"
)

// Server .
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			return handler(ctx, req)
		}
	}
}
