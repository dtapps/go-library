module go.dtapp.net/library/service/chengquan

go 1.24.1

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gorequest => ../../utils/gorequest

require (
	go.dtapp.net/library/utils/gorequest v1.0.96
	go.dtapp.net/library/utils/gotime v1.0.13
)

require github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
