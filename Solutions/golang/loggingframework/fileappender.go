package loggingframework

import (
	"fmt"
	"os"
	"sync"
)

// FileAppender 将日志输出到文件
type FileAppender struct {
	filename string
	file     *os.File
	mu       sync.Mutex // 保护文件写入的互斥锁
}

// NewFileAppender 创建一个新的FileAppender
func NewFileAppender(filename string) (*FileAppender, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	return &FileAppender{
		filename: filename,
		file:     file,
	}, nil
}

// Append 实现LogAppender接口，将日志写入文件
func (f *FileAppender) Append(message *LogMessage) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.file != nil {
		fmt.Fprintln(f.file, message.GetFormattedMessage())
	}
}

// Close 关闭文件
func (f *FileAppender) Close() error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.file != nil {
		err := f.file.Close()
		f.file = nil
		return err
	}
	return nil
}
