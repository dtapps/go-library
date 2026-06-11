module go.dtapp.net/library/service/aswzk

go 1.26.0

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gorequest => ../../utils/gorequest

replace go.dtapp.net/library/utils/gomd5 => ../../utils/gomd5

require (
	go.dtapp.net/library/utils/gomd5 v1.0.3
	go.dtapp.net/library/utils/gorequest v1.1.3
	go.dtapp.net/library/utils/gotime v1.0.19
	resty.dev/v3 v3.0.0-rc.1
)

require (
	github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/text v0.38.0 // indirect
)
