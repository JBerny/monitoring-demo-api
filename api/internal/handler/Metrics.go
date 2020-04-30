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



var metrics = metric.Metrics{
	metric.Metric{
		Counter: promauto.NewCounter(
			prometheus.CounterOpts{
				Help: "Total count of requests",
				Name: "api_request_count", 
			},
		),
		Generator: generator.Rand{ Max: 30, },
	},
	metric.Metric{
		Gauge: promauto.NewGauge(
			prometheus.GaugeOpts{
				Help: "Nr of customers currently logged on",
				Name: "logged_on_customers",
			},
		),
		Generator: generator.NewSin(5*time.Minute, 100),
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
