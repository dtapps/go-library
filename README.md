<h1>
<a href="https://www.dtapp.net/">Golang Library</a>
</h1>

📦 Golang 扩展包

[comment]: <> (go)
[![go](https://github.com/dtapps/go-library/actions/workflows/go.yml/badge.svg)](https://github.com/dtapps/go-library/actions/workflows/go.yml)
[![godoc](https://pkg.go.dev/badge/gopkg.in/dtapps/go-library.v2?status.svg)](https://pkg.go.dev/gopkg.in/dtapps/go-library.v2)
[![goproxy.cn](https://goproxy.cn/stats/gopkg.in/dtapps/go-library.v2/badges/download-count.svg)](https://goproxy.cn/stats/gopkg.in/dtapps/go-library.v2)
[![goreportcard.com](https://goreportcard.com/badge/gopkg.in/dtapps/go-library.v2)](https://goreportcard.com/report/gopkg.in/dtapps/go-library.v2)
[![deps.dev](https://img.shields.io/badge/deps-go-red.svg)](https://deps.dev/go/github.com%2Fdtapps%2Fgo-library)
[![Coverage Status](https://coveralls.io/repos/github/dtapps/go-library/badge.svg?branch=master)](https://coveralls.io/github/dtapps/go-library?branch=master)
[![Sourcegraph](https://sourcegraph.com/github.com/dtapps/go-library/-/badge.svg)](https://sourcegraph.com/github.com/dtapps/go-library?badge)
[![Build status](https://ci.appveyor.com/api/projects/status/d6rq6xynt8wkev5k?svg=true)](https://ci.appveyor.com/project/dtapps/go-library)
[![codecov](https://codecov.io/gh/dtapps/go-library/branch/master/graph/badge.svg?token=BrtbyKKPQX)](https://codecov.io/gh/dtapps/go-library)
[![Build Status](https://app.travis-ci.com/dtapps/go-library.svg?branch=master)](https://app.travis-ci.com/dtapps/go-library)

#### 安装使用

```go
go get -v -u gopkg.in/dtapps/go-library.v2

import (
    "gopkg.in/dtapps/go-library.v2"
)
```

#### 用法示例

> 默认时区为 Local，即服务器所在时区

##### 当前时间

```go
import (
	"gopkg.in/dtapps/go-library.v2/utils/gotime"
)

gotime.Current().Now()
gotime.Current().Format()
gotime.Current().Timestamp()
gotime.Current().TimestampWithMillisecond()
```

## JetBrains 开源证书支持

`go-library` 基于 JetBrains 公司旗下的 GoLand 集成开发环境中进行开发。

<a href="https://www.jetbrains.com/?from=kubeadm-ha" target="_blank">
<img src="https://raw.githubusercontent.com/panjf2000/illustrations/master/jetbrains/jetbrains-variant-4.png" width="250" align="middle"/>
</a>

## 🔑 License

[MIT](https://gopkg.in/dtapps/go-library.v2/blob/master/LICENSE)

Copyright (c) 2018 茂名聚合科技有限公司