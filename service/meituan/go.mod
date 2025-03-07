module go.dtapp.net/library/service/meituan

go 1.23.0

replace go.dtapp.net/library/utils/gojson => ../../utils/gojson

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gostring => ../../utils/gostring

replace go.dtapp.net/library/utils/gorandom => ../../utils/gorandom

replace go.dtapp.net/library/utils/gorequest => ../../utils/gorequest

require (
	go.dtapp.net/library/utils/gorequest v1.0.93
	go.dtapp.net/library/utils/gostring v1.0.24
	go.dtapp.net/library/utils/gotime v1.0.12
)

require (
	github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
	github.com/basgys/goxml2json v1.1.0 // indirect
	go.dtapp.net/library/utils/gojson v1.0.8 // indirect
	go.dtapp.net/library/utils/gorandom v1.0.4 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)
