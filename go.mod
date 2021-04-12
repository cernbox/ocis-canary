module github.com/cernbox/ocis-canary

go 1.13

require (
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	contrib.go.opencensus.io/exporter/ocagent v0.7.0
	contrib.go.opencensus.io/exporter/zipkin v0.1.2
	github.com/UnnoTed/fileb0x v1.1.4
	github.com/asim/go-micro/cmd/protoc-gen-micro/v3 v3.0.0-20210408173139-0d57213d3f5c // indirect
	github.com/asim/go-micro/v3 v3.5.1-0.20210217182006-0f0ace1a44a9
	github.com/cespare/reflex v0.2.0
	github.com/cs3org/reva v1.6.1-0.20210329145723-ed244aac4ddc
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.5.1
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/haya14busa/goverage v0.0.0-20180129164344-eec3514a20b5
	github.com/micro/cli/v2 v2.1.2
	github.com/mitchellh/gox v1.0.1
	github.com/oklog/run v1.1.0
	github.com/openzipkin/zipkin-go v0.2.5
	github.com/owncloud/ocis/ocis-pkg v0.0.0-20210412113235-982264811ecb
	github.com/prometheus/client_golang v1.9.0
	github.com/restic/calens v0.2.0
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/thejerf/suture/v4 v4.0.0
	go.opencensus.io v0.23.0
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
	google.golang.org/genproto v0.0.0-20210207032614-bba0dbe2a9ea
	google.golang.org/protobuf v1.26.0
	honnef.co/go/tools v0.1.1
)

replace (
	// taken from https://github.com/asim/go-micro/blob/master/plugins/registry/etcd/go.mod#L14-L16
	go.etcd.io/etcd/api/v3 => go.etcd.io/etcd/api/v3 v3.0.0-20210204162551-dae29bb719dd
	go.etcd.io/etcd/pkg/v3 => go.etcd.io/etcd/pkg/v3 v3.0.0-20210204162551-dae29bb719dd
	// latest version compatible with etcd
	google.golang.org/grpc => google.golang.org/grpc v1.29.1
)
