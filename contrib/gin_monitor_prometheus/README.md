## 使用

**成功请求的吞吐量 - 折线图**
```
sum(rate(gin_server_throughput{statusCode=~"2.."}[2m])) by (method)
```

**成功请求的 99% 延迟 - 折线图**
```
histogram_quantile(0.99,
  sum(rate(gin_server_latency_us_bucket{statusCode=~"2.."}[2m])) by (method, le)
)
```

**请求失败率 - Stat**
```
(
  sum(rate(gin_server_throughput{statusCode=~"5.."}[2m])) 
  /
  sum(rate(gin_server_throughput[2m]))
) * 100
```

**每秒请求数 - 折线图**
```
sum(rate(gin_server_throughput[1m]))
```

**请求量按状态码拆分 - 柱状图**
```
sum(rate(gin_server_throughput[2m])) by (statusCode)
```