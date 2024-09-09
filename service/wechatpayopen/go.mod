module go.dtapp.net/library/service/wechatpayopen

go 1.23

replace go.dtapp.net/library/utils/gojson => ../../utils/gojson

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gostring => ../../utils/gostring

replace go.dtapp.net/library/utils/gorandom => ../../utils/gorandom

replace go.dtapp.net/library/utils/gorequest => ../../utils/gorequest

require (
	go.dtapp.net/library/utils/gojson v1.0.7
	go.dtapp.net/library/utils/gorandom v1.0.4
	go.dtapp.net/library/utils/gorequest v1.0.83
	go.opentelemetry.io/otel v1.29.0
	go.opentelemetry.io/otel/trace v1.29.0
)

require (
	github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.dtapp.net/library/utils/gostring v1.0.21 // indirect
	go.dtapp.net/library/utils/gotime v1.0.12 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.54.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.54.0 // indirect
	go.opentelemetry.io/otel/metric v1.29.0 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/text v0.18.0 // indirect
)
