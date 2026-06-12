module go.dtapp.net/library/utils/gojobs

go 1.26.0

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gorequest => ../../utils/gorequest

require (
	entgo.io/ent v0.14.6
	github.com/google/uuid v1.6.0
	github.com/redis/go-redis/v9 v9.20.1
	go.dtapp.net/library/utils/gorequest v1.1.3
	go.dtapp.net/library/utils/gotime v1.0.19
	golang.org/x/sync v0.21.0
)

require (
	github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/text v0.38.0 // indirect
)
