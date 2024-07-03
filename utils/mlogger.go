package utils

import (
	go_colorable "github.com/mattn/go-colorable"
	"log"
	"os"
	"runtime"
)

// 定义日志级别
const (
	LogLevelError = iota
	LogLevelSuccess
	LogLevelInfo
)

// 定义颜色代码
const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorReset  = "\033[0m"
)

// 定义日志结构体
type Logger struct {
	logger *log.Logger
	level  int
}

var Mlogger *Logger

// 创建新的日志实例
func NewLogger(level int) *Logger {
	var out *os.File
	if runtime.GOOS == "windows" {
		out = os.Stdout
		return &Logger{
			logger: log.New(go_colorable.NewColorable(out), "", log.LstdFlags),
			level:  level,
		}
	} else {
		out = os.Stdout
		return &Logger{
			logger: log.New(out, "", log.LstdFlags),
			level:  level,
		}
	}
}

// 输出错误日志
func (l *Logger) Error(format string, v ...interface{}) {
	if l.level >= LogLevelError {
		l.logger.Printf(ColorRed+"[ERROR] "+format+ColorReset, v...)
	}
}

// 输出成功日志
func (l *Logger) Success(format string, v ...interface{}) {
	if l.level >= LogLevelSuccess {
		l.logger.Printf(ColorGreen+"[SUCCESS] "+format+ColorReset, v...)
	}
}

// 输出信息日志
func (l *Logger) Info(format string, v ...interface{}) {
	if l.level >= LogLevelInfo {
		l.logger.Printf(ColorYellow+"[INFO] "+format+ColorReset, v...)
	}
}

func main() {
	// 创建一个日志级别为 INFO 的日志实例
	logger := NewLogger(LogLevelInfo)

	// 输出不同级别的日志
	logger.Error("This is an error message: %v", "something went wrong")
	logger.Success("Operation completed successfully!")
	logger.Info("This is an informational message.")
}
