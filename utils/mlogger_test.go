package utils

import (
	"testing"
)

func Test_main(t *testing.T) {
	// 创建一个日志级别为 INFO 的日志实例
	logger := NewLogger(LogLevelInfo)
	logger.level = LogLevelError
	// 输出不同级别的日志
	logger.Error("This is an error message: %v", "something went wrong")
	logger.Success("Operation completed successfully!")
	logger.Info("This is an informational message.")
}
