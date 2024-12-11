<h1>
<a href="https://www.dtapp.net/">Golang Library</a>
</h1>

📦 Golang 扩展包

[comment]: <> (go)
[![godoc](https://pkg.go.dev/badge/github.com/dtapps/go-library?status.svg)](https://pkg.go.dev/github.com/dtapps/go-library)
[![goproxy.cn](https://goproxy.cn/stats/github.com/dtapps/go-library/badges/download-count.svg)](https://goproxy.cn/stats/github.com/dtapps/go-library)
[![goreportcard.com](https://goreportcard.com/badge/github.com/dtapps/go-library)](https://goreportcard.com/report/github.com/dtapps/go-library)
[![deps.dev](https://img.shields.io/badge/deps-go-red.svg)](https://deps.dev/go/github.com%2Fdtapps%2Fgo-library)

#### 用法示例

> 默认时区为 Local，即服务器所在时区

##### 当前时间

```go
import (
	"go.dtapp.net/library/utils/gotime"
)

gotime.Current().Now()
gotime.Current().Format()
gotime.Current().Timestamp()
gotime.Current().TimestampWithMillisecond()
```

## JetBrains 开源证书支持

`go-library` 基于 JetBrains 公司旗下的 GoLand 集成开发环境中进行开发。

<a href="https://www.jetbrains.com/?from=go-library" target="_blank">
<img src="https://user-images.githubusercontent.com/16093106/208018478-27f4911e-9543-4799-a595-8fe4a91b182b.png" width="250" align="middle"/>
</a>

## 🔑 License

[MIT](https://github.com/dtapps/go-library/blob/master/LICENSE)

Copyright (c) 2018 李光春