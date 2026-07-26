# http_log

HTTP 请求/响应日志拦截器，作为 `http.RoundTripper` 中间件使用，自动记录每次 HTTP 调用的完整信息。

## 功能特性

- 自动捕获请求方法、URL、请求头、请求体、状态码、响应头、响应体
- 记录请求耗时（`ElapseTime`）和处理耗时（`ProcessElapseTime`）
- 通过 `httptrace` 捕获网线上真正发送的请求头（含 Transport 层自动添加的头）
- 智能识别 Body 类型（JSON / XML / form-urlencoded / 文本），统一输出为 JSON
- 支持 gzip 压缩内容的自动解压
- 支持接口（`LogHandler`）和回调（`LogCallback`）两种日志处理方式
- 支持调试模式，便于排查问题
- 支持克隆实例，可灵活控制是否记录 Body

## 快速开始

### 方式一：回调模式

```go
package main

import (
    "context"
    "fmt"
    "net/http"

    http_log "github.com/your-org/go-library/contrib/http_log"
)

func main() {
    // 创建拦截器
    logger := http_log.NewLoggingRoundTripper(
        nil,    // Transport，nil 表示使用 http.DefaultTransport
        nil,    // LogHandler 接口
        func(ctx context.Context, data *http_log.LogData) error {
            // 处理日志：写入数据库、发送到消息队列等
            fmt.Printf("[%s] %s %s -> %d (%dms)\n",
                data.Method, data.URL, data.StatusCode, data.ElapseTime)
            return nil
        },
    )

    // 创建 HTTP 客户端
    client := &http.Client{Transport: logger}

    // 发送请求（自动记录日志）
    resp, err := client.Get("https://api.example.com/users")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}
```

### 方式二：接口模式

```go
type MyLogSaver struct{}

func (s *MyLogSaver) HandleLog(ctx context.Context, data *http_log.LogData) error {
    // 持久化日志数据
    return nil
}

func main() {
    saver := &MyLogSaver{}
    logger := http_log.NewLoggingRoundTripper(nil, saver, nil)

    client := &http.Client{Transport: logger}
    // 使用 client 发送请求...
}
```

### 方式三：中间件模式

```go
baseLogger := http_log.NewLoggingRoundTripper(nil, saver, nil)

// 作为中间件链中的一环
middleware := baseLogger.Middleware()
client := &http.Client{Transport: middleware(http.DefaultTransport)}
```

## API 参考

### 核心类型

```go
type LogData struct {
    GoVersion         string              // Go 版本
    PluginVersion     string              // 插件版本
    Hostname          string              // 主机名
    Method            string              // 请求方法
    URL               string              // 请求 URL
    RequestHeaders    map[string][]string // 请求头
    RequestBody       json.RawMessage     // 请求体（JSON 格式）
    StatusCode        int                 // 状态码
    ResponseHeaders   map[string][]string // 响应头
    ResponseBody      json.RawMessage     // 响应体（JSON 格式）
    ElapseTime        int64               // 请求耗时（毫秒）
    ProcessElapseTime int64               // 处理耗时（毫秒）
    IsError           bool                // 是否出错（状态码 >= 400 或读取失败）
}

type LogHandler interface {
    HandleLog(ctx context.Context, data *LogData) error
}

type LogCallback func(ctx context.Context, data *LogData) error
```

### 构造函数

```go
func NewLoggingRoundTripper(
    base http.RoundTripper,  // 底层 Transport，nil 使用 http.DefaultTransport
    handler LogHandler,      // 接口方式处理日志
    callback LogCallback,    // 回调方式处理日志（优先级高于 handler）
) *LoggingRoundTripper
```

### 方法

| 方法 | 说明 |
|------|------|
| `EnableDebug()` | 开启调试模式，打印详细日志到 stdout |
| `Clone()` | 复制一个相同配置的实例 |
| `CloneNoBody()` | 复制一个禁用 Body 记录的实例 |
| `Middleware()` | 返回中间件函数，用于链式组合 |
| `MiddlewareNoBody()` | 返回禁用 Body 记录的中间件函数 |
| `Instance(next)` | 返回配置好的 RoundTripper 实例 |
| `InstanceNoBody(next)` | 返回禁用 Body 记录的 RoundTripper 实例 |

## Body 处理逻辑

根据 `Content-Type` 和内容特征自动识别 Body 类型，统一输出为 `json.RawMessage`：

| Content-Type | 内容特征 | 处理方式 |
|---|---|---|
| `application/json` | 合法 JSON | 原样保留 |
| `application/json` | 不合法 JSON | 降级为 JSON 字符串 |
| `application/xml` / `text/xml` | 合法 XML | 转换为 JSON 对象 |
| `text/html` | 以 `<` 开头 | XML 校验失败 → 降级为 JSON 字符串 |
| `application/x-www-form-urlencoded` | key=value 格式 | 解析为 JSON 对象 |
| `text/plain` | 纯文本 | 作为 JSON 字符串 |
| `multipart/form-data` | 含边界符 | 作为 JSON 字符串 |
| `application/octet-stream` | 二进制 | 作为 JSON 字符串 |
| 空 body | 无内容 | 返回 `nil` |

### 降级策略

当 JSON 或 XML 校验失败时，自动降级为文本处理（作为 JSON 字符串返回），确保**任何内容都不会丢失**。

### gzip 解压

如果 `Content-Encoding: gzip` 存在，会先解压再处理 Body。解压失败则使用原始数据。

## 请求头捕获机制

使用 `httptrace.ClientTrace` 的 `WroteHeaderField` 回调，捕获网线上真正发送的请求头。这包括：

- 应用层通过 `req.Header.Set()` 设置的头部
- Transport 层自动添加的头部（如 `User-Agent`、`Accept-Encoding`）
- 过滤掉 HTTP/2 伪头部（`:method`、`:path` 等）

## 日志输出方式

日志输出采用**异步**方式（goroutine），不阻塞请求流程：

- 如果设置了 `OnLog` 回调，优先使用回调
- 否则使用 `Handler` 接口
- 如果两者都未设置，不记录日志（仅拦截不输出）
- 输出失败时仅打印到 stdout，不影响请求

## 配置选项

```go
logger := http_log.NewLoggingRoundTripper(nil, saver, nil)

// 开启调试模式
logger.EnableDebug()

// 克隆一个不记录 Body 的实例（适合文件上传/下载场景）
noBodyLogger := logger.CloneNoBody()
client := &http.Client{Transport: noBodyLogger}
```

## 注意事项

1. **Body 读取**：请求体和响应体会被完整读入内存，超大 Body 可能影响内存使用
2. **重定向**：已正确处理 `req.GetBody`，跟随重定向时不会丢失请求体
3. **错误处理**：Body 读取失败会返回错误并标记 `IsError = true`
4. **并发安全**：`LogData` 在拦截器内部创建，每个请求独立，无并发竞争
5. **Transport 选择**：`nil` 时使用 `http.DefaultTransport`，建议生产环境传入自定义 Transport 以配置连接池、超时等

## 示例：生产环境用法

```go
transport := &http.Transport{
    MaxIdleConns:        100,
    MaxIdleConnsPerHost: 10,
    IdleConnTimeout:     90 * time.Second,
}

logger := http_log.NewLoggingRoundTripper(transport, &MyLogSaver{}, nil)

client := &http.Client{
    Transport: logger,
    Timeout:   30 * time.Second,
}
```

## 版本历史

| 版本 | 说明 |
|------|------|
| 1.0.13 | 初始版本 |
| 本次修复 | 1) 修复非 JSON/XML 请求体丢失（form-urlencoded 等）<br>2) 修复 ElapseTime 计算错误<br>3) 修复重定向丢失请求体<br>4) 修复 ReadAll 错误被忽略<br>5) JSON/XML 校验失败降级为文本处理 |
