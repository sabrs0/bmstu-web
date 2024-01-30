package benchmark

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	Timings *prometheus.SummaryVec
	Counter *prometheus.CounterVec
}

func NewMetrics() *Metrics {
	return &Metrics{
		Timings: prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Name: "method_timing",
				Help: "per method timing",
			},
			[]string{"method"},
		),
		Counter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "method_counter",
				Help: "Per method counter",
			},
			[]string{"method"},
		),
	}

}
