module go.dtapp.net/library/service/wechatpayopen

go 1.24.1

replace go.dtapp.net/library/utils/gorandom => ../../utils/gorandom

replace go.dtapp.net/library/utils/gorequest => ../../utils/gorequest

require (
	go.dtapp.net/library/utils/gorandom v1.0.4
	go.dtapp.net/library/utils/gorequest v1.0.95
)

require github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
