module go.dtapp.net/library/service/tencent

go 1.25.1

replace go.dtapp.net/library/contrib/resty_log => ../../contrib/resty_log

require (
	go.dtapp.net/library/contrib/resty_log v1.0.8
	resty.dev/v3 v3.0.0-beta.4
)

require golang.org/x/net v0.47.0 // indirect
