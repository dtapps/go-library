module go.dtapp.net/library/utils/gojobs

go 1.23.0

replace go.dtapp.net/library/utils/gojson => ../../utils/gojson

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gorequest => ../../utils/gorequest

replace go.dtapp.net/library/utils/gostring => ../../utils/gostring

replace go.dtapp.net/library/utils/gorandom => ../../utils/gorandom

require (
	entgo.io/ent v0.14.1
	github.com/redis/go-redis/v9 v9.7.0
	github.com/robfig/cron/v3 v3.0.1
	go.dtapp.net/library/utils/gojson v1.0.8
	go.dtapp.net/library/utils/gorequest v1.0.91
	go.dtapp.net/library/utils/gostring v1.0.21
	go.dtapp.net/library/utils/gotime v1.0.12
	golang.org/x/sync v0.9.0
	gorm.io/gorm v1.25.12
)

require (
	github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	go.dtapp.net/library/utils/gorandom v1.0.4 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/text v0.20.0 // indirect
)
