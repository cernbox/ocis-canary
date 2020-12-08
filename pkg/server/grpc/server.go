package grpc

import (
	"github.com/cernbox/ocis-canary/pkg/proto/v0"
	svc "github.com/cernbox/ocis-canary/pkg/service/v0"
	"github.com/cernbox/ocis-canary/pkg/version"
	"github.com/owncloud/ocis/ocis-pkg/service/grpc"
)

// Server initializes the grpc service and server.
func Server(opts ...Option) grpc.Service {
	options := newOptions(opts...)

	service := grpc.NewService(
		grpc.Logger(options.Logger),
		grpc.Name(options.Name),
		grpc.Version(version.String),
		grpc.Address(options.Config.GRPC.Addr),
		grpc.Namespace(options.Config.GRPC.Namespace),
		grpc.Context(options.Context),
		grpc.Flags(options.Flags...),
	)

	handler := svc.NewService(options.Config, options.Logger)
	handler = svc.NewInstrument(handler, options.Metrics)
	handler = svc.NewLogging(handler, options.Logger)
	if err := proto.RegisterCanaryHandler(service.Server(), handler); err != nil {
		options.Logger.Fatal().Err(err).Msg("could not register Canary service handler")
	}

	service.Init()
	return service
}
