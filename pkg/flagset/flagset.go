package flagset

import (
	"github.com/cernbox/ocis-canary/pkg/config"
	"github.com/micro/cli/v2"
)

// RootWithConfig applies cfg to the root flagset
func RootWithConfig(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "config-file",
			Value:       "",
			Usage:       "Path to config file",
			EnvVars:     []string{"CANARY_CONFIG_FILE"},
			Destination: &cfg.File,
		},
		&cli.StringFlag{
			Name:        "log-level",
			Value:       "info",
			Usage:       "Set logging level",
			EnvVars:     []string{"CANARY_LOG_LEVEL"},
			Destination: &cfg.Log.Level,
		},
		&cli.BoolFlag{
			Name:        "log-pretty",
			Usage:       "Enable pretty logging",
			EnvVars:     []string{"CANARY_LOG_PRETTY"},
			Destination: &cfg.Log.Pretty,
		},
		&cli.BoolFlag{
			Name:        "log-color",
			Usage:       "Enable colored logging",
			EnvVars:     []string{"CANARY_LOG_COLOR"},
			Destination: &cfg.Log.Color,
		},
	}
}

// HealthWithConfig applies cfg to the root flagset
func HealthWithConfig(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "debug-addr",
			Value:       "0.0.0.0:9109",
			Usage:       "Address to debug endpoint",
			EnvVars:     []string{"CANARY_DEBUG_ADDR"},
			Destination: &cfg.Debug.Addr,
		},
	}
}

// ServerWithConfig applies cfg to the root flagset
func ServerWithConfig(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "tracing-enabled",
			Usage:       "Enable sending traces",
			EnvVars:     []string{"CANARY_TRACING_ENABLED"},
			Destination: &cfg.Tracing.Enabled,
		},
		&cli.StringFlag{
			Name:        "tracing-type",
			Value:       "jaeger",
			Usage:       "Tracing backend type",
			EnvVars:     []string{"CANARY_TRACING_TYPE"},
			Destination: &cfg.Tracing.Type,
		},
		&cli.StringFlag{
			Name:        "tracing-endpoint",
			Value:       "",
			Usage:       "Endpoint for the agent",
			EnvVars:     []string{"CANARY_TRACING_ENDPOINT"},
			Destination: &cfg.Tracing.Endpoint,
		},
		&cli.StringFlag{
			Name:        "tracing-collector",
			Value:       "",
			Usage:       "Endpoint for the collector",
			EnvVars:     []string{"CANARY_TRACING_COLLECTOR"},
			Destination: &cfg.Tracing.Collector,
		},
		&cli.StringFlag{
			Name:        "tracing-service",
			Value:       "canary",
			Usage:       "Service name for tracing",
			EnvVars:     []string{"CANARY_TRACING_SERVICE"},
			Destination: &cfg.Tracing.Service,
		},
		&cli.StringFlag{
			Name:        "debug-addr",
			Value:       "0.0.0.0:9109",
			Usage:       "Address to bind debug server",
			EnvVars:     []string{"CANARY_DEBUG_ADDR"},
			Destination: &cfg.Debug.Addr,
		},
		&cli.StringFlag{
			Name:        "debug-token",
			Value:       "",
			Usage:       "Token to grant metrics access",
			EnvVars:     []string{"CANARY_DEBUG_TOKEN"},
			Destination: &cfg.Debug.Token,
		},
		&cli.BoolFlag{
			Name:        "debug-pprof",
			Usage:       "Enable pprof debugging",
			EnvVars:     []string{"CANARY_DEBUG_PPROF"},
			Destination: &cfg.Debug.Pprof,
		},
		&cli.BoolFlag{
			Name:        "debug-zpages",
			Usage:       "Enable zpages debugging",
			EnvVars:     []string{"CANARY_DEBUG_ZPAGES"},
			Destination: &cfg.Debug.Zpages,
		},
		&cli.StringFlag{
			Name:        "http-namespace",
			Value:       "ch.cernbox.web",
			Usage:       "Set the base namespace for the http namespace",
			EnvVars:     []string{"CANARY_HTTP_NAMESPACE"},
			Destination: &cfg.HTTP.Namespace,
		},
		&cli.StringFlag{
			Name:        "http-addr",
			Value:       "0.0.0.0:9105",
			Usage:       "Address to bind http server",
			EnvVars:     []string{"CANARY_HTTP_ADDR"},
			Destination: &cfg.HTTP.Addr,
		},
		&cli.StringFlag{
			Name:        "http-root",
			Value:       "/",
			Usage:       "Root path of http server",
			EnvVars:     []string{"CANARY_HTTP_ROOT"},
			Destination: &cfg.HTTP.Root,
		},
		&cli.IntFlag{
			Name:        "http-cache-ttl",
			Value:       604800, // 7 days
			Usage:       "Set the static assets caching duration in seconds",
			EnvVars:     []string{"CANARY_CACHE_TTL"},
			Destination: &cfg.HTTP.CacheTTL,
		},
		&cli.StringFlag{
			Name:        "grpc-namespace",
			Value:       "ch.cernbox.api",
			Usage:       "Set the base namespace for the grpc namespace",
			EnvVars:     []string{"CANARY_GRPC_NAMESPACE"},
			Destination: &cfg.GRPC.Namespace,
		},
		&cli.StringFlag{
			Name:        "grpc-addr",
			Value:       "0.0.0.0:9106",
			Usage:       "Address to bind grpc server",
			EnvVars:     []string{"CANARY_GRPC_ADDR"},
			Destination: &cfg.GRPC.Addr,
		},
		&cli.StringFlag{
			Name:        "asset-path",
			Value:       "",
			Usage:       "Path to custom assets",
			EnvVars:     []string{"CANARY_ASSET_PATH"},
			Destination: &cfg.Asset.Path,
		},
		&cli.StringFlag{
			Name:        "jwt-secret",
			Value:       "Pive-Fumkiu4",
			Usage:       "Used to create JWT to talk to reva, should equal reva's jwt-secret",
			EnvVars:     []string{"CANARY_JWT_SECRET"},
			Destination: &cfg.TokenManager.JWTSecret,
		},
		&cli.StringFlag{
			Name:        "db-host",
			Value:       "localhost",
			Usage:       "Host of the mysql DB",
			EnvVars:     []string{"CANARY_DB_HOST"},
			Destination: &cfg.Service.DB.Host,
		},
		&cli.IntFlag{
			Name:        "db-port",
			Value:       3306,
			Usage:       "Port of the mysql DB",
			EnvVars:     []string{"CANARY_DB_PORT"},
			Destination: &cfg.Service.DB.Port,
		},
		&cli.StringFlag{
			Name:        "db-name",
			Value:       "cernbox",
			Usage:       "Name of the mysql DB",
			EnvVars:     []string{"CANARY_DB_NAME"},
			Destination: &cfg.Service.DB.Name,
		},
		&cli.StringFlag{
			Name:        "db-username",
			Value:       "",
			Usage:       "Username of the write account for mysql DB",
			EnvVars:     []string{"CANARY_DB_USERNAME"},
			Destination: &cfg.Service.DB.Username,
		},
		&cli.StringFlag{
			Name:        "db-password",
			Value:       "",
			Usage:       "Password of the write account for mysql DB",
			EnvVars:     []string{"CANARY_DB_PASSWORD"},
			Destination: &cfg.Service.DB.Password,
		},
		&cli.StringFlag{
			Name:        "db-table",
			Value:       "cbox_canary",
			Usage:       "Mysql DB table name for storing the canary info",
			EnvVars:     []string{"CANARY_DB_TABLE"},
			Destination: &cfg.Service.DB.Table,
		},
		&cli.IntFlag{
			Name:        "cookie-ttl",
			Value:       10440,
			Usage:       "TTL for the Canary cookie",
			EnvVars:     []string{"CANARY_COOKIE_TTL"},
			Destination: &cfg.Service.Cookie.TTL,
		},
		&cli.StringFlag{
			Name:        "cookie-name",
			Value:       "web_canary",
			Usage:       "Name for the Canary cookie",
			EnvVars:     []string{"CANARY_COOKIE_Name"},
			Destination: &cfg.Service.Cookie.Name,
		},
	}
}
