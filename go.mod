module github.com/cernbox/ocis-canary

go 1.13

require (
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	contrib.go.opencensus.io/exporter/ocagent v0.7.0
	contrib.go.opencensus.io/exporter/zipkin v0.1.1
	github.com/UnnoTed/fileb0x v1.1.4
	github.com/cespare/reflex v0.2.0
	github.com/cs3org/reva v1.4.1-0.20201203135805-655c9babdc3f
	github.com/davecgh/go-spew v1.1.1
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/haya14busa/goverage v0.0.0-20180129164344-eec3514a20b5
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/kr/pty v1.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/protoc-gen-micro/v2 v2.1.1 // indirect
	github.com/mitchellh/gox v1.0.1
	github.com/ogier/pflag v0.0.1 // indirect
	github.com/oklog/run v1.1.0
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/owncloud/ocis v1.0.0-rc6 // indirect
	github.com/owncloud/ocis/ocis-pkg v0.0.0-20201208093523-507ad7ba9168
	github.com/prometheus/client_golang v1.7.1
	github.com/restic/calens v0.2.0
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.6.1
	go.opencensus.io v0.22.5
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	google.golang.org/genproto v0.0.0-20200624020401-64a14ca9d1ad
	google.golang.org/protobuf v1.25.0
	honnef.co/go/tools v0.0.1-2020.1.5
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
