package engine

import (
	"time"

	"github.com/jberny/monitoring-demo-api/pkg/generator"
	"github.com/jberny/monitoring-demo-api/pkg/metric"
	"github.com/prometheus/client_golang/prometheus"
)

// Opts contains options to configure the engine.Run
type Opts struct {
	Interval  time.Duration
	Metric    metric.Metric
}

// Run executes a Task with configured interval
func Run(o Opts) {
	go func() {
		for {
			if o.Metric.Counter != nil {
				counterAdd(o.Metric.Generator, o.Metric.Counter)
			} else if o.Metric.Gauge != nil {
				gaugeSet(o.Metric.Generator, o.Metric.Gauge)
			} else {
				return
			}
			time.Sleep(o.Interval)
		}
	}()
}

func counterAdd(g generator.Generator, m prometheus.Counter)  {
	m.Add(g.NextVal())
	
}

func gaugeSet(g generator.Generator, m prometheus.Gauge)  {
	m.Set(g.NextVal())
}