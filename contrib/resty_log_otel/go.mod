module go.dtapp.net/library/contrib/resty_log_otel

go 1.25.1

replace go.dtapp.net/library/contrib/resty_log => ../../contrib/resty_log

require (
	go.dtapp.net/library/contrib/resty_log v1.0.6
	go.opentelemetry.io/otel v1.38.0
	go.opentelemetry.io/otel/trace v1.38.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel/metric v1.38.0 // indirect
	golang.org/x/net v0.46.0 // indirect
	resty.dev/v3 v3.0.0-beta.3 // indirect
)
