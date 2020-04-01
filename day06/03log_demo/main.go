package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// log demo
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

func main() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	log.SetOutput(logFile)
	for {
		fmt.Println("这是一条测试的日志")
		log.Println("这是一条测试的日志")
		time.Sleep(time.Second)
	}
}
