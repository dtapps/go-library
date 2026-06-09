module go.dtapp.net/library/utils/gojobs

go 1.24.2

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gorequest => ../../utils/gorequest

require (
	entgo.io/ent v0.14.6
	github.com/google/uuid v1.6.0
	github.com/redis/go-redis/v9 v9.7.3
	go.dtapp.net/library/utils/gorequest v1.0.93
	go.dtapp.net/library/utils/gotime v1.0.13
	golang.org/x/sync v0.17.0
)

require (
	github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	golang.org/x/net v0.46.0 // indirect
	golang.org/x/text v0.30.0 // indirect
)
