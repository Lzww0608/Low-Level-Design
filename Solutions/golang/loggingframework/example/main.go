package main

import (
	"loggingframework"
)

func main() {
	// 创建一个Logger，最小日志级别为DEBUG
	logger := loggingframework.NewLogger("AppLogger", loggingframework.LogLevelDebug)

	// 添加控制台输出器
	consoleAppender := loggingframework.NewConsoleAppender()
	logger.AddAppender(consoleAppender)

	// 添加文件输出器
	fileAppender, err := loggingframework.NewFileAppender("app.log")
	if err != nil {
		panic(err)
	}
	defer fileAppender.Close()
	logger.AddAppender(fileAppender)

	// 使用不同级别记录日志
	logger.Debug("This is a debug message", "main")
	logger.Info("Application started", "main")
	logger.Warning("This is a warning", "main")
	logger.Error("An error occurred", "main")
	logger.Fatal("Fatal error!", "main")

	// 创建一个只记录WARNING及以上级别的Logger
	warnLogger := loggingframework.NewLogger("WarnLogger", loggingframework.LogLevelWarning)
	warnLogger.AddAppender(consoleAppender)

	// 这条不会被记录（级别太低）
	warnLogger.Debug("This debug won't be logged", "main")
	warnLogger.Info("This info won't be logged", "main")

	// 这些会被记录
	warnLogger.Warning("This warning will be logged", "main")
	warnLogger.Error("This error will be logged", "main")
}

