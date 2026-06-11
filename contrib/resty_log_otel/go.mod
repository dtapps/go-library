module go.dtapp.net/library/contrib/resty_log_otel

go 1.26.0

replace go.dtapp.net/library/contrib/resty_log => ../../contrib/resty_log

require (
	go.dtapp.net/library/contrib/resty_log v1.0.15
	go.opentelemetry.io/otel v1.44.0
	go.opentelemetry.io/otel/trace v1.44.0
	resty.dev/v3 v3.0.0-rc.1
)

require (
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel/metric v1.44.0 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/text v0.38.0 // indirect
)
