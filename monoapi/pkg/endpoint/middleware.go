package endpoint

import (
	"context"
	"fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	metrics "github.com/go-kit/kit/metrics"
	"go.uber.org/zap"
	"time"
)

// InstrumentingMiddleware returns an endpoint middleware that records
// the duration of each invocation to the passed histogram. The middleware adds
// a single field: "success", which is "true" if no error is returned, and
// "false" otherwise.
func InstrumentingMiddleware(duration metrics.Histogram) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				duration.With("success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}

// LoggingMiddleware returns an endpoint middleware that logs the
// duration of each invocation, and the resulting error, if any.
func LoggingMiddleware(logger *zap.SugaredLogger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				if err != nil {
					logger.Error("transport_error", err, "took", time.Since(begin))
				} else {
					logger.Info("Done ", "took: ", time.Since(begin))
				}
			}(time.Now())
			return next(ctx, request)
		}
	}
}
