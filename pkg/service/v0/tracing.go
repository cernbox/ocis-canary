package svc

import (
	"context"

	v0proto "github.com/cernbox/ocis-canary/pkg/proto/v0/github.com/cernbox/ocis-canary/pkg/proto/v0"
	"go.opencensus.io/trace"
)

// NewTracing returns a service that instruments traces.
func NewTracing(next v0proto.CanaryHandler) v0proto.CanaryHandler {
	return tracing{
		next: next,
	}
}

type tracing struct {
	next v0proto.CanaryHandler
}

// SetCanary implements the CanaryHandler interface.
func (t tracing) SetCanary(ctx context.Context, req *v0proto.CanaryRequest, rsp *v0proto.CanaryResponse) error {
	ctx, span := trace.StartSpan(ctx, "Canary.SetCanary")
	defer span.End()

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("name", req.Version),
	}, "Execute Canary.SetCanary handler")

	return t.next.SetCanary(ctx, req, rsp)
}
