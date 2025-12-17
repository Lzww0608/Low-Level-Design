package loggingframework

import (
	"sync"
	"time"
)

// Logger 是日志记录器的主要类
type Logger struct {
	name      string
	minLevel  LogLevel
	appenders []LogAppender
	mu        sync.RWMutex // 保护appenders的读写锁
}

// NewLogger 创建一个新的Logger实例
// name: 日志器的名称
// minLevel: 最小日志级别，默认为INFO
func NewLogger(name string, minLevel LogLevel) *Logger {
	return &Logger{
		name:      name,
		minLevel:  minLevel,
		appenders: make([]LogAppender, 0),
	}
}

// AddAppender 添加一个日志输出器
func (l *Logger) AddAppender(appender LogAppender) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.appenders = append(l.appenders, appender)
}

// Log 记录指定级别的日志
func (l *Logger) Log(level LogLevel, message, source string) {
	if l.isLevelEnabled(level) {
		logMessage := NewLogMessage(time.Now(), level, message, source)

		l.mu.RLock()
		defer l.mu.RUnlock()

		for _, appender := range l.appenders {
			appender.Append(logMessage)
		}
	}
}

// Debug 记录DEBUG级别的日志
func (l *Logger) Debug(message, source string) {
	l.Log(LogLevelDebug, message, source)
}

// Info 记录INFO级别的日志
func (l *Logger) Info(message, source string) {
	l.Log(LogLevelInfo, message, source)
}

// Warning 记录WARNING级别的日志
func (l *Logger) Warning(message, source string) {
	l.Log(LogLevelWarning, message, source)
}

// Error 记录ERROR级别的日志
func (l *Logger) Error(message, source string) {
	l.Log(LogLevelError, message, source)
}

// Fatal 记录FATAL级别的日志
func (l *Logger) Fatal(message, source string) {
	l.Log(LogLevelFatal, message, source)
}

// isLevelEnabled 检查给定的日志级别是否启用
func (l *Logger) isLevelEnabled(level LogLevel) bool {
	return level >= l.minLevel
}

// GetName 返回日志器的名称
func (l *Logger) GetName() string {
	return l.name
}

// SetMinLevel 设置最小日志级别
func (l *Logger) SetMinLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.minLevel = level
}

// GetMinLevel 获取最小日志级别
func (l *Logger) GetMinLevel() LogLevel {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.minLevel
}
