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
				Objectives: map[float64]float64{ 0.99: 0.1, 0.95: 0.2, 0.5: 0.3 },
			},
			[]string{"url"},
		),
		Generator: generator.APIRequestDuration{},
		Labels: map[string][]string{
			"url": {"/api", "/api/users", "/api/books/naked-sun", 
				"/api/authors/isaac-asimov", 
				"/api/authors/isaac-asimov/books",
			},
		},
	},
	metric.Metric{
		Histogram: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Help: "Total time in seconds that it invoke a service and receive a response",
				Name: "service_request_duration_seconds",
				Buckets: []float64{ 0.1, 0.2 },
			},
			[]string{"service"},
		),
		Generator: generator.ServiceRequestDuration{},
		Labels: map[string][]string{
			"service": {"users", "authors", "books", },
		},
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

