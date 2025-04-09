module go.dtapp.net/library/utils/gophp

go 1.23.0

replace go.dtapp.net/library/utils/gostring => ../../utils/gostring

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gorandom => ../../utils/gorandom

require go.dtapp.net/library/utils/gostring v1.0.25

require (
	go.dtapp.net/library/utils/gorandom v1.0.4 // indirect
	go.dtapp.net/library/utils/gotime v1.0.13 // indirect
)
