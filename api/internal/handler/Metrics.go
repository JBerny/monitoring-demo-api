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
	"country": []string{ "nl", "it", "de", "pl", "fr" },
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
