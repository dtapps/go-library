module go.dtapp.net/library/utils/gostring

go 1.23

replace go.dtapp.net/library/utils/gojson => ../../utils/gojson

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gorandom => ../../utils/gorandom

require (
	go.dtapp.net/library/utils/gojson v1.0.8
	go.dtapp.net/library/utils/gorandom v1.0.4
	go.dtapp.net/library/utils/gotime v1.0.12
)

require (
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
