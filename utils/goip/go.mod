module go.dtapp.net/library/utils/goip

go 1.25.0

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gostring => ../../utils/gostring

replace go.dtapp.net/library/utils/gorandom => ../../utils/gorandom

replace go.dtapp.net/library/utils/gorequest => ../../utils/gorequest

require (
	github.com/ip2location/ip2location-go/v9 v9.7.0
	github.com/lionsoul2014/ip2region/binding/golang v0.0.0-20260408140759-61ef15cff068
	github.com/oschwald/geoip2-golang/v2 v2.1.0
	github.com/tagphi/czdb-search-golang v1.0.4
	go.dtapp.net/library/utils/gorequest v1.1.3
)

require (
	github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/oschwald/maxminddb-golang/v2 v2.1.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/net v0.52.0 // indirect
	golang.org/x/sys v0.43.0 // indirect
	golang.org/x/text v0.35.0 // indirect
	lukechampine.com/uint128 v1.3.0 // indirect
)
