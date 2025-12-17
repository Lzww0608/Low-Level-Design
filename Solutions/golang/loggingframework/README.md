# Go 日志框架 (Logging Framework)

这是一个从C++移植到Go的简单但功能完整的日志框架实现。

## 功能特性

- ✅ 支持多个日志级别：DEBUG, INFO, WARNING, ERROR, FATAL
- ✅ 可配置的最小日志级别过滤
- ✅ 支持多个输出目标（Appenders）
- ✅ 控制台输出（ConsoleAppender）
- ✅ 文件输出（FileAppender）
- ✅ 线程安全（使用互斥锁保护）
- ✅ 灵活的日志格式化

## 架构设计

### 核心组件

1. **LogLevel** (`loglevel.go`)
   - 定义日志级别常量
   - 提供级别到字符串的转换

2. **LogMessage** (`logmessage.go`)
   - 封装日志消息的所有信息
   - 包含时间戳、级别、消息内容和来源

3. **LogAppender** (`logappender.go`)
   - 定义输出器接口
   - 所有输出器必须实现`Append`方法

4. **ConsoleAppender** (`consoleappender.go`)
   - 将日志输出到标准输出

5. **FileAppender** (`fileappender.go`)
   - 将日志输出到文件
   - 支持追加模式
   - 线程安全的文件写入

6. **Logger** (`logger.go`)
   - 核心日志记录器
   - 管理多个Appenders
   - 提供便捷的日志方法

## 使用示例

### 基本使用

```go
package main

import "loggingframework"

func main() {
    // 创建Logger
    logger := loggingframework.NewLogger("MyApp", loggingframework.LogLevelInfo)
    
    // 添加控制台输出
    logger.AddAppender(loggingframework.NewConsoleAppender())
    
    // 记录日志
    logger.Info("Application started", "main")
    logger.Warning("This is a warning", "main")
    logger.Error("An error occurred", "main")
}
```

### 同时输出到控制台和文件

```go
logger := loggingframework.NewLogger("MyApp", loggingframework.LogLevelDebug)

// 添加控制台输出
logger.AddAppender(loggingframework.NewConsoleAppender())

// 添加文件输出
fileAppender, err := loggingframework.NewFileAppender("app.log")
if err != nil {
    panic(err)
}
defer fileAppender.Close()
logger.AddAppender(fileAppender)

// 所有日志会同时输出到控制台和文件
logger.Debug("Debug info", "module")
logger.Info("Info message", "module")
```

### 使用不同的日志级别

```go
// 只记录WARNING及以上级别
warnLogger := loggingframework.NewLogger("WarnOnly", loggingframework.LogLevelWarning)
warnLogger.AddAppender(loggingframework.NewConsoleAppender())

warnLogger.Debug("Won't be logged", "test")  // 不会输出
warnLogger.Info("Won't be logged", "test")   // 不会输出
warnLogger.Warning("Will be logged", "test") // 会输出
warnLogger.Error("Will be logged", "test")   // 会输出
```

### 动态修改日志级别

```go
logger := loggingframework.NewLogger("DynamicLogger", loggingframework.LogLevelInfo)

// 运行时修改日志级别
logger.SetMinLevel(loggingframework.LogLevelDebug)

// 获取当前日志级别
currentLevel := logger.GetMinLevel()
```

## 运行示例程序

```bash
cd example
go run main.go
```

## 与C++版本的对应关系

| C++ | Go |
|-----|-----|
| `Logger::Logger()` | `NewLogger()` |
| `Logger::addAppender()` | `Logger.AddAppender()` |
| `Logger::log()` | `Logger.Log()` |
| `Logger::debug()` | `Logger.Debug()` |
| `Logger::info()` | `Logger.Info()` |
| `Logger::warning()` | `Logger.Warning()` |
| `Logger::error()` | `Logger.Error()` |
| `Logger::fatal()` | `Logger.Fatal()` |
| `std::shared_ptr<LogAppender>` | `LogAppender` interface |
| `ConsoleAppender` | `ConsoleAppender` |
| `FileAppender` | `FileAppender` |

## 线程安全

- `Logger` 使用 `sync.RWMutex` 保护appenders列表的并发访问
- `FileAppender` 使用 `sync.Mutex` 保护文件写入操作
- 所有公共方法都是线程安全的

## 扩展性

你可以通过实现 `LogAppender` 接口来创建自定义的输出器：

```go
type MyCustomAppender struct {
    // 你的字段
}

func (m *MyCustomAppender) Append(message *loggingframework.LogMessage) {
    // 你的实现
    // 例如：输出到数据库、网络、消息队列等
}
```

## 日志格式

默认的日志格式为：
```
[LEVEL] SOURCE: MESSAGE
```

例如：
```
[INFO] main: Application started
[ERROR] database: Connection failed
```

## 文件结构

```
loggingframework/
├── go.mod
├── loglevel.go          # 日志级别定义
├── logmessage.go        # 日志消息结构
├── logappender.go       # 输出器接口
├── consoleappender.go   # 控制台输出器
├── fileappender.go      # 文件输出器
├── logger.go            # 主日志器
├── README.md            # 本文档
└── example/
    └── main.go          # 使用示例
```

## 许可证

本项目是学习和练习低层设计（LLD）的一部分。

