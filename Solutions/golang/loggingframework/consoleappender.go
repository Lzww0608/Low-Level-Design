package loggingframework

import "fmt"

// ConsoleAppender 将日志输出到控制台
type ConsoleAppender struct{}

// NewConsoleAppender 创建一个新的ConsoleAppender
func NewConsoleAppender() *ConsoleAppender {
	return &ConsoleAppender{}
}

// Append 实现LogAppender接口，将日志输出到控制台
func (c *ConsoleAppender) Append(message *LogMessage) {
	fmt.Println(message.GetFormattedMessage())
}
