package handler

import (
	"net/http"
	"time"

	"github.com/jberny/monitoring-demo-api/pkg/engine"
	"github.com/jberny/monitoring-demo-api/pkg/metric"
	"github.com/jberny/monitoring-demo-api/pkg/generator"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var labels = []string{"country"}

var mapLabels = map[string][]string{
	"country": { "nl", "it", "de", "pl", "fr" },
}

var metrics = metric.Metrics{
	metric.Metric{
		Counters: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Help: "Total count of requests",
				Name: "api_request_count", 
			},
			labels,
		),
		Generator: generator.Rand{ Max: 30, },
		Labels: mapLabels,
	},
	metric.Metric{
		Gauges: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Help: "Nr of customers currently logged on",
				Name: "logged_on_customers",
			},
			labels,
		),
		Generator: generator.NewLoggedOnCustomers(),
		Labels: mapLabels,
	},
	metric.Metric{
		Summary: promauto.NewSummaryVec(
			prometheus.SummaryOpts{
				Help: "Total time in seconds that it takes to the api to fulfill a request",
				Name: "api_request_duration_seconds",
				Objectives: map[float64]float64{ 0.99: 0.0, 0.95: 0.0, 0.5: 0.0 },
			},
			labels,
		),
		Generator: generator.APIRequestDuration{},
		Labels: mapLabels,
	},
	metric.Metric{
		Histogram: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Help: "Total time in seconds that it invoke a service and receive a response",
				Name: "service_request_duration_seconds",
				Buckets: []float64{ 100.0, 200.0 },
			},
			labels,
		),
		Generator: generator.ServiceRequestDuration{},
		Labels: mapLabels,
	},
}

// NewMetrics initializes a new Metric struct
func NewMetrics() http.Handler {
	for _, metric := range metrics {
		opts := engine.Opts{
			Interval: 1*time.Second,
			Metric: metric,
		}
		engine.Run(opts)
	}
	return promhttp.Handler()
}

// func (m Metrics) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	m.Handler.ServeHTTP(w, r)
// }
