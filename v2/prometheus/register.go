package prometheus

import (
	"github.com/Azure/go-shuttle/v2/prometheus/listener"
	"github.com/Azure/go-shuttle/v2/prometheus/publisher"
	"github.com/prometheus/client_golang/prometheus"
)

func Register(registerer prometheus.Registerer) {
	listener.Metrics.Init(registerer)
	publisher.Metrics.Init(registerer)
}
