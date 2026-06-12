module go.dtapp.net/library/utils/goip-czdb

go 1.26.0

replace go.dtapp.net/library/utils/gojson => ../../utils/gojson

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gostring => ../../utils/gostring

replace go.dtapp.net/library/utils/gorandom => ../../utils/gorandom

require (
	github.com/zhengjianyang/goCzdb v0.1.4
	go.dtapp.net/library/utils/gostring v1.0.32
)

require (
	github.com/spf13/cast v1.10.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	go.dtapp.net/library/utils/gorandom v1.0.5 // indirect
	go.dtapp.net/library/utils/gotime v1.0.19 // indirect
)
