package svc

import (
	"context"
	"time"

	v0proto "github.com/cernbox/ocis-canary/pkg/proto/v0"
	"github.com/owncloud/ocis/ocis-pkg/log"
)

// NewLogging returns a service that logs messages.
func NewLogging(next v0proto.CanaryHandler, logger log.Logger) v0proto.CanaryHandler {
	return logging{
		next:   next,
		logger: logger,
	}
}

type logging struct {
	next   v0proto.CanaryHandler
	logger log.Logger
}

// SetCanary implements the CanaryHandler interface.
func (l logging) SetCanary(ctx context.Context, req *v0proto.CanaryRequest, rsp *v0proto.CanaryResponse) error {
	start := time.Now()
	err := l.next.SetCanary(ctx, req, rsp)

	logger := l.logger.With().
		Str("method", "Canary.SetCanary").
		Dur("duration", time.Since(start)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("Failed to execute")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}
