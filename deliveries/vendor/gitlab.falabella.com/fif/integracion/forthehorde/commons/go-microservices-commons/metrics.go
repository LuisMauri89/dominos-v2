package commons

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

//MetricsConfig configuracion de metricas para endpoints
type MetricsConfig struct {
	RequestDuration    metrics.Histogram
	SuccessfulRequests metrics.Counter
	FailedRequests     metrics.Counter
}

//MakeDefaultEndpointMetrics crea metricas por defecto para endpoint middleware de metricas
func MakeDefaultEndpointMetrics(namespace, subsystem string) *MetricsConfig {
	return &MetricsConfig{
		RequestDuration: prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "request_duration_seconds",
			Help:      "Duracion de request en segundos.",
		}, []string{"operation", "success"}),
		SuccessfulRequests: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "successful_requests",
			Help:      "Total de peticiones exitosas.",
		}, []string{}),
		FailedRequests: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "failed_requests",
			Help:      "Total de peticiones fallidas.",
		}, []string{}),
	}
}
