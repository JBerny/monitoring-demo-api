package metric

import (
	"github.com/jberny/monitoring-demo-api/pkg/generator"
	"github.com/prometheus/client_golang/prometheus"
)

// Metric represents a prometheus metric
type Metric struct {
	Counter     prometheus.Counter
	Gauge 	    prometheus.Gauge
	Generator   generator.Generator
}

// Metrics an array of Metric
type Metrics []Metric

