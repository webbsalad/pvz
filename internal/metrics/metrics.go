package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	PVZCreated       = prometheus.NewCounter(prometheus.CounterOpts{Name: "pvz_created_total", Help: "how many pvz created "})
	ReceptionCreated = prometheus.NewCounter(prometheus.CounterOpts{Name: "reception_created_total", Help: "how many receptions created"})
	ProductAdded     = prometheus.NewCounter(prometheus.CounterOpts{Name: "product_added_total", Help: "how many products added"})
)

func init() {
	prometheus.MustRegister(
		PVZCreated,
		ReceptionCreated,
		ProductAdded,
	)
}
