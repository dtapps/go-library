# req_log

基于 [req](https://github.com/imroc/req) v3 的 HTTP 请求/响应日志中间件库，自动捕获请求参数、响应体、耗时等关键信息，支持回调和接口两种日志输出方式。

## 特性

- 自动捕获请求方法、URL、Header、Body、Query/Form/Path 参数
- 自动捕获响应状态码、Header、Body
- 智能识别 JSON / XML / form-urlencoded / 纯文本等 Content-Type
- 支持请求耗时（ElapseTime）和处理耗时（ProcessElapseTime）
- 支持回调函数（OnLog）和接口（Handler）两种日志输出方式
- 支持 `Clone` / `CloneNoBody` 快速复制配置
- 异步日志发送，不阻塞请求链路
- 请求 Context 取消后仍可发送日志

## 安装

```bash
go get go.dtapp.net/library/contrib/req_log
```

## 快速开始

### 方式一：回调函数

```go
package main

import (
    "context"
    "fmt"
    "go.dtapp.net/library/contrib/req_log"
    "github.com/imroc/req/v3"
)

func main() {
    client := req.C()

    // 创建中间件，使用回调函数
    logger := req_log.NewLoggerMiddleware(nil, func(ctx context.Context, data *req_log.LogData) error {
        fmt.Printf("[%s] %s %d -> %d (%dms)\n",
            data.Method, data.URL, data.StatusCode, data.ElapseTime)
        return nil
    })

    // 注册中间件
    client.OnBeforeRequest(logger.OnBeforeRequest)
    client.OnAfterResponse(logger.OnAfterResponse)

    // 发起请求
    resp, err := client.R().
        SetBody(map[string]string{"name": "alice"}).
        SetHeader("Content-Type", "application/json").
        Post("https://httpbin.org/post")
}
```

### 方式二：接口实现

```go
// 实现 LogHandler 接口
type MyLogSaver struct{}

func (s *MyLogSaver) HandleLog(ctx context.Context, data *req_log.LogData) error {
    // 保存到数据库/消息队列/文件等
    jsonData, _ := json.Marshal(data)
    fmt.Println(string(jsonData))
    return nil
}

func main() {
    client := req.C()
    logger := req_log.NewLoggerMiddleware(&MyLogSaver{}, nil)

    client.OnBeforeRequest(logger.OnBeforeRequest)
    client.OnAfterResponse(logger.OnAfterResponse)
}
```

### 方式三：测试辅助

```go
// 使用内置的测试 Logger（同时使用回调和接口，回调优先）
logger := req_log.NewTestLogger()
```

## 中间件注册

```go
client := req.C()

// 1. OnBeforeRequest — 注入开始时间，记录请求信息
client.OnBeforeRequest(logger.OnBeforeRequest)

// 2. OnAfterResponse — 汇总日志数据并触发保存
client.OnAfterResponse(logger.OnAfterResponse)
```

## API 参考

### 构造函数

```go
func NewLoggerMiddleware(handler LogHandler, callback LogCallback) *LoggerMiddleware
```

- `handler` — 实现 `LogHandler` 接口的对象
- `callback` — 回调函数 `func(ctx context.Context, data *LogData) error`
- 两者都传时，**callback 优先**

### 方法

| 方法 | 说明 |
|------|------|
| `EnableDebug()` | 开启调试模式，打印详细日志 |
| `Clone()` | 复制一个新的 LoggerMiddleware（保留所有配置） |
| `CloneNoBody()` | 复制一个新的 LoggerMiddleware，强制禁用请求体和响应体记录 |

### Hook 方法

| Hook | 说明 |
|------|------|
| `OnBeforeRequest` | 注入开始时间到 Context |
| `OnAfterResponse` | 汇总日志数据并触发保存 |

### LogData 结构

```go
type LogData struct {
    GoVersion          string              // Go 版本
    PluginVersion      string              // 插件版本
    Hostname           string              // 主机名
    Method             string              // 请求方法
    URL                string              // 请求 URL
    RequestHeaders     map[string][]string // 请求头
    RequestBody        json.RawMessage     // 请求体
    RequestQueryParams map[string][]string // 请求参数
    RequestFormData    map[string][]string // 请求表单数据
    RequestPathParams  map[string]string   // 请求路径参数
    StatusCode         int                 // 状态码
    ResponseHeaders    map[string][]string // 响应头
    ResponseBody       json.RawMessage     // 响应体
    ElapseTime         int64               // 耗时（毫秒）
    ProcessElapseTime  int64               // 处理耗时（毫秒）
    IsError            bool                // 是否错误（状态码 >= 400）
}
```

### Context Key 函数

```go
func WithStartTimeKey(ctx context.Context, t time.Time) context.Context
func GetStartTimeKey(ctx context.Context) time.Time

func WithRequestBodyKey(ctx context.Context, body []byte) context.Context
func GetRequestBodyKey(ctx context.Context) []byte

func WithResponseBodyKey(ctx context.Context, body []byte) context.Context
func GetResponseBodyKey(ctx context.Context) []byte
```

## Body 处理逻辑

### 类型检测优先级

1. **Content-Type 匹配**：`application/json` → JSON，`xml`/`soap+xml` → XML
2. **内容嗅探**：`{`/`[` 开头 → JSON，`<` 开头 → XML
3. **格式校验**：尝试 JSON/XML 解析
4. **兜底**：纯文本

### 处理结果

| 类型 | 处理方式 |
|------|----------|
| JSON | 校验通过直接返回；失败降级为 JSON 字符串 |
| XML | 校验通过转换为 JSON；失败降级为 JSON 字符串 |
| form-urlencoded | 解析为 JSON 对象（单值→string，多值→[]string） |
| 其他文本 | 包装为 JSON 字符串 |
| 二进制 | 包装为 JSON 字符串 |
| 空 | 返回 `nil` |

### 常量定义

```go
const (
    JSONBodyType = "json"
    XMLBodyType  = "xml"
    TextBodyType = "text"

    ContentTypeHeader = "Content-Type"

    ApplicationJSON = "application/json"
    ApplicationXML  = "application/xml"
    TextXML         = "text/xml"
    SOAPXML         = "soap+xml"
    FormURLEncoded  = "application/x-www-form-urlencoded"
)
```

## 性能优化

- `detectBodyType` 使用字节比较（`data[0]`）避免 `string(data)` 内存分配
- XML 转换使用 `bytes.NewReader` 零拷贝
- `context.WithoutCancel` 统一在 `emit` 内部处理，避免重复调用
- 异步日志发送，不阻塞请求链路
- Goroutine panic 恢复，防止用户回调 panic 导致应用崩溃

## 版本历史

| 版本 | 说明 |
|------|------|
| v1.0.7 | 初始版本 |
| v1.0.8 | 修复 ElapseTime 计算错误、添加 TextBodyType 处理、JSON/XML 降级策略、性能优化、panic 恢复 |
