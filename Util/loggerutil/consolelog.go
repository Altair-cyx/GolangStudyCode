package loggerutil

import (
	"fmt"
	"time"
)

// 1.支持往不同的地方输出日志

// 2.日志级别
// logger.Trace()
// logger.Debug()
// logger.Warning()
// logger.Info()
// logger.Error()
// logger.Fatal()

// 3.日志要支持开关控制，比如说开发的时候什么级别都能输出，但是上线之后只有INFO级别往下的才能输出

// 4.完整的日志记录要包含有时间、行号、文件名、日志级别、日志信息

// 5.日志文件要切割
// 		1.按文件大小切割
//			1.每次记录日志之前都判断一下当前写的这个文件的大小
//		2.按日期切割
//			1.在日志结构体重设置一个字段记录上一次切割的小时数
//			2.在写日志之前检查一下当前时间的小时数和之前保存的是否一致，不一致就切割

// ConsoleLog 控制台日志结构体
type ConsoleLog struct{}

// NewConsoleLog 构造函数
func NewConsoleLog() ConsoleLog {
	return ConsoleLog{}
}

// All ...
func (c ConsoleLog) All(format string, a ...interface{}) {
	logTime := time.Now()
	// 完整的日志记录要包含有时间、行号、文件名、日志级别、日志信息
	fmt.Printf("%s %s\n", logTime.Format("2006-01-02 15:04:05.000"), format)
}

// Trace ...
func (c ConsoleLog) Trace(msg string) {
	fmt.Println(msg, time.Now())
}

// Debug ...
func (c ConsoleLog) Debug(msg string) {
	fmt.Println(msg, time.Now())
}

// Info ...
func (c ConsoleLog) Info(msg string) {
	fmt.Println(msg, time.Now())
}

// Warn ...
func (c ConsoleLog) Warn(msg string) {
	fmt.Println(msg, time.Now())
}

// Error ...
func (c ConsoleLog) Error(msg string) {
	fmt.Println(msg, time.Now())
}

// Fatal ...
func (c ConsoleLog) Fatal(msg string) {
	fmt.Println(msg, time.Now())
}

// Off ...
func (c ConsoleLog) Off(msg string) {
	fmt.Println(msg, time.Now())
}
