module github.com/cernbox/ocis-canary

go 1.13

require (
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	contrib.go.opencensus.io/exporter/ocagent v0.7.0
	contrib.go.opencensus.io/exporter/zipkin v0.1.2
	github.com/UnnoTed/fileb0x v1.1.4
	github.com/asim/go-micro/v3 v3.5.1-0.20210217182006-0f0ace1a44a9
	github.com/cespare/reflex v0.2.0
	github.com/cs3org/reva v1.10.1-0.20210730095301-fcb7a30a44a6
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/haya14busa/goverage v0.0.0-20180129164344-eec3514a20b5
	github.com/micro/cli/v2 v2.1.2
	github.com/mitchellh/gox v1.0.1
	github.com/oklog/run v1.1.0
	github.com/openzipkin/zipkin-go v0.2.5
	github.com/owncloud/ocis/ocis-pkg v0.0.0-20210216094451-dc73176dc62d
	github.com/prometheus/client_golang v1.10.0
	github.com/restic/calens v0.2.0
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/thejerf/suture/v4 v4.0.1
	go.opencensus.io v0.23.0
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	google.golang.org/genproto v0.0.0-20210624195500-8bfb893ecb84
	google.golang.org/grpc/examples v0.0.0-20210803221256-6ba56c814be7 // indirect
	google.golang.org/protobuf v1.27.1
	honnef.co/go/tools v0.1.1
)

replace (
	github.com/owncloud/ocis => github.com/cernbox/ocis v0.0.0-20210804114859-0317150cd4e8
	github.com/owncloud/ocis/ocis-pkg => github.com/cernbox/ocis/ocis-pkg v0.0.0-20210804114859-0317150cd4e8
	go.etcd.io/etcd/api/v3 => go.etcd.io/etcd/api/v3 v3.0.0-20210204162551-dae29bb719dd
	go.etcd.io/etcd/pkg/v3 => go.etcd.io/etcd/pkg/v3 v3.0.0-20210204162551-dae29bb719dd
)
