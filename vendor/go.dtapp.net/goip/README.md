<h1>
<a href="https://www.dtapp.net/">Golang Ip</a>
</h1>

📦 Golang Ip库组件

[comment]: <> (go)
[![godoc](https://pkg.go.dev/badge/go.dtapp.net/goip?status.svg)](https://pkg.go.dev/go.dtapp.net/goip)
[![goproxy.cn](https://goproxy.cn/stats/go.dtapp.net/goip/badges/download-count.svg)](https://goproxy.cn/stats/go.dtapp.net/goip)
[![goreportcard.com](https://goreportcard.com/badge/go.dtapp.net/goip	)](https://goreportcard.com/report/go.dtapp.net/goip)
[![deps.dev](https://img.shields.io/badge/deps-go-red.svg)](https://deps.dev/go/go.dtapp.net/goip)

#### 安装使用

```go
go get -v -u go.dtapp.net/goip
```

#### 导入

```go
import (
    "go.dtapp.net/goip"
)
```

```go
package main

import (
	"go.dtapp.net/goip"
	"testing"
)

func TestGoIp(t *testing.T) {
	// 获取Mac地址
	t.Log(goip.GetMacAddr())
	// 内网ip
	t.Log(goip.GetInsideIp())
	// 外网ip
	t.Log(goip.GetOutsideIp())
}

```