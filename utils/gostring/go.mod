module go.dtapp.net/library/utils/gostring

go 1.22.5

replace go.dtapp.net/library/utils/gojson => ../../utils/gojson

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gorandom => ../../utils/gorandom

require (
	go.dtapp.net/library/utils/gojson v1.0.7
	go.dtapp.net/library/utils/gorandom v1.0.4
	go.dtapp.net/library/utils/gotime v1.0.11
)

require (
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
