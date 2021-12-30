package golog

import (
	"log"
	"testing"
)

var app = App{
	LogPath:      "./",
	LogName:      "all.log",
	LogLevel:     "debug",
	MaxSize:      2,
	MaxBackups:   30,
	MaxAge:       0,
	Compress:     false,
	JsonFormat:   false,
	ShowLine:     true,
	LogInConsole: true,
}

func TestLog(t *testing.T) {
	log.Println(app.Logger)
	app.InitLogger()
	app.Logger.Debug("debug 日志")
	app.Logger.Sugar().Debug("debug 日志")
	app.Logger.Info("info 日志")
	app.Logger.Sugar().Info("info 日志")
	app.Logger.Warn("warning 日志")
	app.Logger.Sugar().Warn("warning 日志")
	app.Logger.Error("error 日志")
	app.Logger.Sugar().Error("error 日志")
	log.Println(app.Logger)
}
