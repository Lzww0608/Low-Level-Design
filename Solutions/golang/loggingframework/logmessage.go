package loggingframework

import (
	"fmt"
	"time"
)

type LogMessage struct {
	Timestamp time.Time
	Level     LogLevel
	Message   string
	Source    string
}

func NewLogMessage(timestamp time.Time, level LogLevel, message, source string) *LogMessage {
	return &LogMessage{
		Timestamp: timestamp,
		Level:     level,
		Message:   message,
		Source:    source,
	}
}

func (l *LogMessage) GetTimestamp() time.Time {
	return l.Timestamp
}

func (l *LogMessage) GetLevel() LogLevel {
	return l.Level
}

func (l *LogMessage) GetMessage() string {
	return l.Message
}

func (l *LogMessage) GetSource() string {
	return l.Source
}

func (l *LogMessage) GetFormattedMessage() string {
	return fmt.Sprintf("[%s] %s: %s", logLevelToString(l.Level), l.Source, l.Message)
}

func (l *LogMessage) SetTimestamp(timestamp time.Time) {
	l.Timestamp = timestamp
}

func (l *LogMessage) SetLevel(level LogLevel) {
	l.Level = level
}

func (l *LogMessage) SetMessage(message string) {
	l.Message = message
}

func (l *LogMessage) SetSource(source string) {
	l.Source = source
}
