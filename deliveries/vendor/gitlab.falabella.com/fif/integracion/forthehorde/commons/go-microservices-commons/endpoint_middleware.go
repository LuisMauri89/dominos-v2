package commons

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/metrics"
)

//MakeDefaultEntryEndpoint endpoint middleware por defecto
func MakeDefaultEntryEndpoint(service string, logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return EndpointLogMiddleware(service, logger)(next)
	}
}

// EndpointLogMiddleware loguea tiempo que tomo servir request y error si es que hubo alguno
func EndpointLogMiddleware(service string, logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		logger = log.WithPrefix(logger, "caller", log.DefaultCaller)
		logger = log.WithPrefix(logger, "serviceName", service)

		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				level.Debug(logger).Log("started", begin, "took", time.Since(begin))

				if err != nil {
					level.Error(logger).Log("transport_error", err)
				}
			}(time.Now())
			return next(ctx, request)
		}
	}
}

// EndpointSuccesAndFailureMetricsMiddleware middleware de metrics de success y failures
func EndpointSuccesAndFailureMetricsMiddleware(success, failures metrics.Counter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func() {
				if err != nil {
					failures.Add(1)
				} else {
					success.Add(1)
				}
			}()

			return next(ctx, request)
		}
	}
}

// EndpointTimeTakenMetricsMiddleware middleware de metrics para tiempo que tomo la peticion
func EndpointTimeTakenMetricsMiddleware(operation string, duration metrics.Histogram) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				duration.With("operation", operation, "success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}
