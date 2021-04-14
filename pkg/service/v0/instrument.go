package svc

import (
	"context"

	"github.com/cernbox/ocis-canary/pkg/metrics"
	v0proto "github.com/cernbox/ocis-canary/pkg/proto/v0/github.com/ocis-canary"
	"github.com/prometheus/client_golang/prometheus"
)

// NewInstrument returns a service that instruments metrics.
func NewInstrument(next v0proto.CanaryHandler, metrics *metrics.Metrics) v0proto.CanaryHandler {
	return instrument{
		next:    next,
		metrics: metrics,
	}
}

type instrument struct {
	next    v0proto.CanaryHandler
	metrics *metrics.Metrics
}

// SetCanary implements the CanaryHandler interface.
func (i instrument) SetCanary(ctx context.Context, req *v0proto.CanaryRequest, rsp *v0proto.CanaryResponse) error {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
		us := v * 1000000

		i.metrics.Latency.WithLabelValues().Observe(us)
		i.metrics.Duration.WithLabelValues().Observe(v)
	}))

	defer timer.ObserveDuration()

	err := i.next.SetCanary(ctx, req, rsp)

	if err == nil {
		i.metrics.Counter.WithLabelValues().Inc()
	}

	return err
}
