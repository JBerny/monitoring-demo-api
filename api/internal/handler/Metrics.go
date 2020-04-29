package handler

import (
	"net/http"

	"github.com/jberny/monitoring-demo-api/pkg/generator"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	Internal  http.Handler
	Generator generator.Generator
	Custom    prometheus.Gauge
}

// NewMetrics initializes a new Metrics struct
func NewMetrics() *Metrics {
	return &Metrics{
		Internal: promhttp.Handler(),
		Generator: generator.NewSin(10.0),
		Custom: promauto.NewGauge(prometheus.GaugeOpts{
			Help: "Custom example metric, generates a sine of 10Hz",
			Name: "custom_metric",
		}),
	}
}

func (m Metrics) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	val := m.Generator.NextVal()
	m.Custom.Set(val)
	m.Internal.ServeHTTP(w, r)
}
