package engine

import (
	"time"

	"github.com/jberny/monitoring-demo-api/pkg/metric"
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
			for _, v := range o.Metric.Labels {
				for _, lv := range v { 
					val := o.Metric.Generator.NextVal()
					if o.Metric.Counters != nil {
						c, err := o.Metric.Counters.GetMetricWithLabelValues(lv)
						if err == nil {
							c.Add(val)
						}
					} else if o.Metric.Gauges != nil {
						g, err := o.Metric.Gauges.GetMetricWithLabelValues(lv)
						if err == nil {
							g.Set(val)
						}
					} else if o.Metric.Summary != nil {
						s, err := o.Metric.Summary.GetMetricWithLabelValues(lv)
						if err == nil {
							s.Observe(val)
						}
					} else if o.Metric.Histogram != nil {
						s, err := o.Metric.Histogram.GetMetricWithLabelValues(lv)
						if err == nil {
							s.Observe(val)
						}
					} else {
						return
					}
				}
			}
			time.Sleep(o.Interval)
		}
	}()
}
