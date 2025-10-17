module go.dtapp.net/library/service/pinduoduo

go 1.25.1

replace go.dtapp.net/library/contrib/resty_log => ../../contrib/resty_log

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gorequest => ../../utils/gorequest

require (
	github.com/shopspring/decimal v1.4.0
	go.dtapp.net/library/contrib/resty_log v1.0.6
	go.dtapp.net/library/utils/gorequest v1.1.2
	go.dtapp.net/library/utils/gotime v1.0.13
	resty.dev/v3 v3.0.0-beta.3
)

require (
	github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.46.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
