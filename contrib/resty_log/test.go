package resty_log

import (
	"context"
	"encoding/json"
	"fmt"
)

// 实现 LogSaver 接口
type TestEntLogSaver struct{}

func (s *TestEntLogSaver) HandleLog(ctx context.Context, data *LogData) error {
	// json 人性化打印 data
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println("[接口方式]打印日志数据:")
	fmt.Println(string(jsonData))
	fmt.Println("-----------------")
	return nil
}

// 返回一个测试的 Logger
func NewTestLogger() *LoggerMiddleware {

	// 定义回调函数
	callbackFunc := func(ctx context.Context, data *LogData) error {
		// json 人性化打印 data
		jsonData, _ := json.MarshalIndent(data, "", "  ")
		fmt.Println("[回调方式]打印日志数据:")
		fmt.Println(string(jsonData))
		fmt.Println("-----------------")
		return nil
	}

	// 创建 LogSaver 实现
	saver := &TestEntLogSaver{}

	// 初始化 LoggerMiddleware 传入回调和接口（回调优先）
	return NewLoggerMiddleware(saver, callbackFunc)
}
