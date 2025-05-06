module go.dtapp.net/library/service/x7s

go 1.24.1

replace go.dtapp.net/library/utils/gomd5 => ../../utils/gomd5

replace go.dtapp.net/library/utils/gorequest => ../../utils/gorequest

require (
	go.dtapp.net/library/utils/gomd5 v1.0.3
	go.dtapp.net/library/utils/gorequest v1.0.96
)

require github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
