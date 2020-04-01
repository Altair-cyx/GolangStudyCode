package loggerutil

// Logger 日志接口方法
type Logger interface {
	All()
	Trace()
	Debug()
	Info()
	Warn()
	Error()
	Fatal()
	Off()
}

// LogLevel 日志级别
type LogLevel uint16

// 日志等级
const (
	ALL LogLevel = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)
