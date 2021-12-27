<h1><a href="https://www.dtapp.net/">Golang Library</a></h1>

📦 Golang 扩展包

[comment]: <> (go)
[![Release](https://img.shields.io/github/release/dtapps/go-library.svg)](https://github.com/dtapps/go-library/releases/latest)
[![go](https://github.com/dtapps/go-library/actions/workflows/go.yml/badge.svg)](https://github.com/dtapps/go-library/actions/workflows/go.yml)
[![godoc](https://pkg.go.dev/badge/github.com/dtapps/go-library?status.svg)](https://pkg.go.dev/github.com/dtapps/go-library)
[![goproxy.cn](https://goproxy.cn/stats/github.com/dtapps/go-library/badges/download-count.svg)](https://goproxy.cn/stats/github.com/dtapps/go-library)
[![goreportcard.com](https://goreportcard.com/badge/github.com/dtapps/go-library)](https://goreportcard.com/report/github.com/dtapps/go-library)
[![deps.dev](https://img.shields.io/badge/deps-go-red.svg)](https://deps.dev/go/github.com%2Fdtapps%2Fgo-library)
[![Coverage Status](https://coveralls.io/repos/github/dtapps/go-library/badge.svg?branch=master)](https://coveralls.io/github/dtapps/go-library?branch=master)
[![Sourcegraph](https://sourcegraph.com/github.com/dtapps/go-library/-/badge.svg)](https://sourcegraph.com/github.com/dtapps/go-library?badge)
[![Build status](https://ci.appveyor.com/api/projects/status/d6rq6xynt8wkev5k?svg=true)](https://ci.appveyor.com/project/dtapps/go-library)
[![codecov](https://codecov.io/gh/dtapps/go-library/branch/master/graph/badge.svg?token=BrtbyKKPQX)](https://codecov.io/gh/dtapps/go-library)


#### 安装使用

```go
go get -v -u github.com/dtapps/go-library

import (
    "github.com/dtapps/go-library"
)
```

#### 用法示例

> 默认时区为 Local，即服务器所在时区

##### 当前时间

```go
import (
	"github.com/dtapps/go-library/utils/gotime"
)

gotime.Current().Now()
gotime.Current().Format()
gotime.Current().Timestamp()
gotime.Current().TimestampWithMillisecond()
```
