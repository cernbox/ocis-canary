package config

// Log defines the available logging configuration.
type Log struct {
	Level  string
	Pretty bool
	Color  bool
}

// Debug defines the available debug configuration.
type Debug struct {
	Addr   string
	Token  string
	Pprof  bool
	Zpages bool
}

// HTTP defines the available http configuration.
type HTTP struct {
	Addr      string
	Namespace string
	Root      string
	CacheTTL  int
}

// GRPC defines the available grpc configuration.
type GRPC struct {
	Addr      string
	Namespace string
}

// Service provides configuration options for the service
type Service struct {
	DB     DB
	Cookie Cookie
}

// DB configuration for service
type DB struct {
	Host     string
	Port     int
	Name     string
	Username string
	Password string
	Table    string
}

// Cookie configuration for Canary
type Cookie struct {
	Name string
	TTL  int
}

// Tracing defines the available tracing configuration.
type Tracing struct {
	Enabled   bool
	Type      string
	Endpoint  string
	Collector string
	Service   string
}

// Asset defines the available asset configuration.
type Asset struct {
	Path string
}

// TokenManager is the config for using the reva token manager
type TokenManager struct {
	JWTSecret string
}

// Config combines all available configuration parts.
type Config struct {
	File         string
	Service      Service
	Log          Log
	Debug        Debug
	HTTP         HTTP
	GRPC         GRPC
	Tracing      Tracing
	Asset        Asset
	TokenManager TokenManager
}

// New initializes a new configuration with or without defaults.
func New() *Config {
	return &Config{}
}
