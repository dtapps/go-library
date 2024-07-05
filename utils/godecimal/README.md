<h1>
<a href="https://www.dtapp.net/">Golang Decimal</a>
</h1>

📦 Golang 小数点处理

[comment]: <> (go)
[![godoc](https://pkg.go.dev/badge/go.dtapp.net/godecimal?status.svg)](https://pkg.go.dev/go.dtapp.net/godecimal)
[![goproxy.cn](https://goproxy.cn/stats/go.dtapp.net/godecimal/badges/download-count.svg)](https://goproxy.cn/stats/go.dtapp.net/godecimal)
[![goreportcard.com](https://goreportcard.com/badge/go.dtapp.net/godecimal)](https://goreportcard.com/report/go.dtapp.net/godecimal)
[![deps.dev](https://img.shields.io/badge/deps-go-red.svg)](https://deps.dev/go/go.dtapp.net%2Fgodecimal)

#### 安装

```shell
go get -v -u go.dtapp.net/godecimal@v1.0.11
```

#### 使用

```go
package main

import (
	"go.dtapp.net/godecimal"
	"log"
	"reflect"
)

func main() {
	log.Println("加：", godecimal.Float64Add(10, 3), reflect.TypeOf(godecimal.Float64Add(10, 3)))
	log.Println("减", godecimal.Float64Sub(10, 3), reflect.TypeOf(godecimal.Float64Sub(10, 3)))
	log.Println("乘：", godecimal.Float64Mul(10, 3), reflect.TypeOf(godecimal.Float64Mul(10, 3)))
	log.Println("除：", godecimal.Float64Quo(10, 3), reflect.TypeOf(godecimal.Float64Quo(10, 3)))
}
```