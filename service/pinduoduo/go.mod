module go.dtapp.net/library/service/pinduoduo

go 1.23.1

replace go.dtapp.net/library/utils/gojson => ../../utils/gojson

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gostring => ../../utils/gostring

replace go.dtapp.net/library/utils/gorandom => ../../utils/gorandom

replace go.dtapp.net/library/utils/gorequest => ../../utils/gorequest

require (
	github.com/shopspring/decimal v1.4.0
	go.dtapp.net/library/utils/gojson v1.0.7
	go.dtapp.net/library/utils/gorequest v1.0.86
	go.dtapp.net/library/utils/gostring v1.0.21
	go.dtapp.net/library/utils/gotime v1.0.12
)

require (
	github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
	github.com/basgys/goxml2json v1.1.0 // indirect
	go.dtapp.net/library/utils/gorandom v1.0.4 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/text v0.20.0 // indirect
)
