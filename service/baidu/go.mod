module go.dtapp.net/library/service/baidu

go 1.24.2

replace go.dtapp.net/library/utils/gojson => ../../utils/gojson

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gostring => ../../utils/gostring

replace go.dtapp.net/library/utils/gorandom => ../../utils/gorandom

replace go.dtapp.net/library/utils/gorequest => ../../utils/gorequest

require (
	go.dtapp.net/library/utils/gojson v0.0.0-00010101000000-000000000000
	go.dtapp.net/library/utils/gorequest v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/otel/trace v1.28.0
)

require (
	github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
	github.com/basgys/goxml2json v1.1.0 // indirect
	go.opentelemetry.io/otel v1.28.0 // indirect
	golang.org/x/net v0.46.0 // indirect
	golang.org/x/text v0.30.0 // indirect
)
