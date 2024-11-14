module go.dtapp.net/library/utils/goip

go 1.23.0

replace go.dtapp.net/library/utils/gojson => ../../utils/gojson

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gostring => ../../utils/gostring

replace go.dtapp.net/library/utils/gorandom => ../../utils/gorandom

replace go.dtapp.net/library/utils/gorequest => ../../utils/gorequest

require (
	github.com/ip2location/ip2location-go/v9 v9.7.0
	github.com/oschwald/geoip2-golang v1.11.0
	go.dtapp.net/library/utils/gorequest v1.0.86
	go.dtapp.net/library/utils/gostring v1.0.21
	golang.org/x/text v0.20.0
)

require (
	github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/oschwald/maxminddb-golang v1.13.1 // indirect
	go.dtapp.net/library/utils/gojson v1.0.7 // indirect
	go.dtapp.net/library/utils/gorandom v1.0.4 // indirect
	go.dtapp.net/library/utils/gotime v1.0.12 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	lukechampine.com/uint128 v1.3.0 // indirect
)
